package updater

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/config"
	"github.com/tarik0/DexEqualizer/dexpair"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/slices"
	"math/big"
	"strings"
	"sync"
	"time"
)

// The abis.

var multicallerAbi, _ = abi.JSON(strings.NewReader(abis.MulticallerMetaData.ABI))
var routerAbi, _ = abi.JSON(strings.NewReader(abis.IPancakeRouter02ABI))
var factoryAbi, _ = abi.JSON(strings.NewReader(abis.IPancakeFactoryABI))
var pairAbi, _ = abi.JSON(strings.NewReader(abis.PairMetaData.ABI))
var erc20Abi, _ = abi.JSON(strings.NewReader(abis.ERC20MetaData.ABI))

var syncId = pairAbi.Events["Sync"].ID
var transferId = erc20Abi.Events["Transfer"].ID

// findFactories
// 	Finds the factory addresses from the routers.
func (p *PairUpdater) findFactories() error {
	// Check length.
	if len(p.params.Routers.Addresses) > 30 {
		return variables.TooManyRouters
	}

	// Address slice.
	var contractAddresses = make([]common.Address, len(p.params.Routers.Addresses))
	for i, router := range p.params.Routers.Addresses {
		contractAddresses[i] = router
	}

	// Function name slice.
	var functionNames = make([]string, len(p.params.Routers.Addresses))
	for i, _ := range p.params.Routers.Addresses {
		functionNames[i] = "factory"
	}

	// Arguments slice.
	var functionArgs = make([][]interface{}, len(p.params.Routers.Addresses))
	for i, _ := range p.params.Routers.Addresses {
		functionArgs[i] = make([]interface{}, 0)
	}

	// Abi slice.
	var abiList = make([]abi.ABI, len(p.params.Routers.Addresses))
	for i, _ := range p.params.Routers.Addresses {
		abiList[i] = routerAbi
	}

	// Multi call factories..
	_, returnBytes, err := p.multiCall(abiList, contractAddresses, functionNames, functionArgs)
	if err != nil {
		return err
	}

	// Check errors.
	if len(returnBytes[0]) == 0 {
		return variables.EmptyResponse
	}

	p.Factories = make([]common.Address, len(returnBytes))
	p.FactoryToRouter = make(map[common.Address]common.Address)
	p.RouterToFactory = make(map[common.Address]common.Address)

	// Decode the factories.
	for i, addrBytes := range returnBytes {
		p.Factories[i] = common.BytesToAddress(addrBytes)
		p.FactoryToRouter[p.Factories[i]] = contractAddresses[i]
		p.RouterToFactory[contractAddresses[i]] = p.Factories[i]
	}

	return nil
}

// findPairAddresses
//	Finds the pair addresses from the factories. (one call costs 1760 gas)
func (p *PairUpdater) findPairAddresses() error {
	// Skip if no factory found.
	if len(p.Factories) == 0 {
		return variables.NoFactoryFound
	}

	// Calculate sizes.
	factoriesSize := len(p.Factories)
	pairsSize := ((len(p.params.Tokens.Addresses) - 1) * len(p.params.Tokens.Addresses)) / 2
	totalPairsSize := factoriesSize * pairsSize

	// Address slice.
	var contractAddresses = make([]common.Address, totalPairsSize)
	for factoryIndex, pairIndex := 0, 0; factoryIndex < factoriesSize; factoryIndex += 1 {
		for tmp := 0; tmp < pairsSize; tmp += 1 {
			contractAddresses[pairIndex] = p.Factories[factoryIndex]
			pairIndex += 1
		}
	}

	// Function name slice.
	var functionNames = make([]string, totalPairsSize)
	for i := 0; i < totalPairsSize; i += 1 {
		functionNames[i] = "getPair"
	}

	// Arguments slice.
	var pairTokens = make([][]common.Address, totalPairsSize)
	var functionArgs = make([][]interface{}, totalPairsSize)
	for factoryIndex, pairIndex := 0, 0; factoryIndex < factoriesSize; factoryIndex += 1 {
		for m := 0; m < len(p.params.Tokens.Addresses); m++ {
			for n := m + 1; n < len(p.params.Tokens.Addresses); n++ {
				functionArgs[pairIndex] = make([]interface{}, 2)
				functionArgs[pairIndex][0] = p.params.Tokens.Addresses[m]
				functionArgs[pairIndex][1] = p.params.Tokens.Addresses[n]

				// Sort tokens.
				pairTokens[pairIndex] = dexpair.SortTokens(p.params.Tokens.Addresses[m], p.params.Tokens.Addresses[n])
				pairIndex += 1
			}
		}
	}

	// Abi slice.
	var abiList = make([]abi.ABI, totalPairsSize)
	for i := 0; i < totalPairsSize; i++ {
		abiList[i] = factoryAbi
	}

	// Split into chunks if too many.
	chunkSize := 300
	var contractAddressesChunks [][]common.Address = chunkBy(contractAddresses, chunkSize)
	var functionNamesChunks [][]string = chunkBy(functionNames, chunkSize)
	var functionArgsChunks [][][]interface{} = chunkBy(functionArgs, chunkSize)

	// Check chunk number.
	if len(contractAddressesChunks) > MaxProcessAmount-1 {
		return variables.TooManyPairs
	}

	// Wait group.
	var wg sync.WaitGroup

	// The response channel.
	var resChan = make(chan *struct {
		Factories     []common.Address
		PairAddresses []common.Address
		PairTokens    [][]common.Address
		Err           error
	})

	// Start workers.
	for chunkId := 0; chunkId < len(contractAddressesChunks); chunkId++ {
		// The chunks.
		contractAddressesChunk := contractAddressesChunks[chunkId]
		functionNamesChunk := functionNamesChunks[chunkId]
		functionArgsChunk := functionArgsChunks[chunkId]

		// Add one to the group.
		wg.Add(1)

		// Start routine.
		go func(addresses []common.Address, funcNames []string, funcArgs [][]interface{}) {
			defer wg.Done()

			// Multi call factories..
			_, returnBytes, err := p.multiCall(abiList, addresses, funcNames, funcArgs)
			if err != nil {
				resChan <- &struct {
					Factories     []common.Address
					PairAddresses []common.Address
					PairTokens    [][]common.Address
					Err           error
				}{nil, nil, nil, err}
				return
			}

			// Check errors.
			if len(returnBytes[0]) == 0 {
				resChan <- &struct {
					Factories     []common.Address
					PairAddresses []common.Address
					PairTokens    [][]common.Address
					Err           error
				}{nil, nil, nil, variables.EmptyResponse}
				return
			}

			// The response variables.
			factories := make([]common.Address, 0)
			pairAddresses := make([]common.Address, 0)
			_pairTokens := make([][]common.Address, 0)

			// Decode pairs.
			for k, addrBytes := range returnBytes {
				// Decode.
				pairAddr := common.BytesToAddress(addrBytes)
				tokenA := (funcArgs[k][0]).(common.Address)
				tokenB := (funcArgs[k][1]).(common.Address)
				factory := addresses[k]

				factories = append(factories, factory)
				pairAddresses = append(pairAddresses, pairAddr)
				_pairTokens = append(_pairTokens, dexpair.SortTokens(tokenA, tokenB))
			}

			resChan <- &struct {
				Factories     []common.Address
				PairAddresses []common.Address
				PairTokens    [][]common.Address
				Err           error
			}{factories, pairAddresses, _pairTokens, nil}
		}(contractAddressesChunk, functionNamesChunk, functionArgsChunk)
	}

	// Wait.
	go func() {
		wg.Wait()
		close(resChan)
	}()

	p.TokenToPairs = make(map[common.Address][]*dexpair.DexPair)
	p.AddressToPair = make(map[common.Address]*dexpair.DexPair)
	p.PairToTokens = make(map[common.Address][]common.Address)
	p.PairToFactory = make(map[common.Address]common.Address)
	p.Pairs = make([]*dexpair.DexPair, 0)

	// Collect results.
	for r := range resChan {
		if r.Err != nil {
			return r.Err
		}

		// Iterate over pairs.
		for i, pairAddr := range r.PairAddresses {
			// Skip empty.
			if pairAddr.String() == "0x0000000000000000000000000000000000000000" {
				continue
			}

			// Get tokens.
			tokenA, tokenB := r.PairTokens[i][0], r.PairTokens[i][1]

			// Create array if not found.
			if val, ok := p.TokenToPairs[tokenA]; !ok || val == nil {
				p.TokenToPairs[tokenA] = make([]*dexpair.DexPair, 0)
			}
			if val, ok := p.TokenToPairs[tokenB]; !ok || val == nil {
				p.TokenToPairs[tokenB] = make([]*dexpair.DexPair, 0)
			}

			// Append to map.
			tmpPair := dexpair.NewDexPair(pairAddr, tokenA, tokenB)
			p.TokenToPairs[tokenA] = append(p.TokenToPairs[tokenA], tmpPair)
			p.TokenToPairs[tokenB] = append(p.TokenToPairs[tokenB], tmpPair)
			p.AddressToPair[pairAddr] = tmpPair
			p.PairToTokens[pairAddr] = r.PairTokens[i]
			p.PairToFactory[pairAddr] = r.Factories[i]
			p.Pairs = append(p.Pairs, tmpPair)
		}
	}

	return nil
}

// findReserves
//	Finds the pair reserves from the pairs.
func (p *PairUpdater) findReserves() error {
	// Skip if no pair.
	if len(p.Pairs) == 0 {
		return variables.NoPairFound
	}

	// Address slice.
	var contractAddresses = make([]common.Address, len(p.Pairs))
	for i, _pair := range p.Pairs {
		contractAddresses[i] = _pair.Address
	}

	// Function name slice.
	var functionNames = make([]string, len(p.Pairs))
	for i, _ := range p.Pairs {
		functionNames[i] = "getReserves"
	}

	// Arguments slice.
	var functionArgs = make([][]interface{}, len(p.Pairs))
	for i, _ := range p.Pairs {
		functionArgs[i] = make([]interface{}, 0)
	}

	// Abi slice.
	var abiList = make([]abi.ABI, len(p.Pairs))
	for i, _ := range p.Pairs {
		abiList[i] = pairAbi
	}

	// Split into chunks if too many.
	chunkSize := 300
	var contractAddressesChunks [][]common.Address = chunkBy(contractAddresses, chunkSize)
	var functionNamesChunks [][]string = chunkBy(functionNames, chunkSize)
	var functionArgsChunks [][][]interface{} = chunkBy(functionArgs, chunkSize)

	// Check chunk number.
	if len(contractAddressesChunks) > MaxProcessAmount-1 {
		return variables.TooManyPairs
	}

	// Wait group.
	var wg sync.WaitGroup

	// The response channel.
	var resChan = make(chan *struct {
		Reserves  [][]*big.Int
		Addresses []common.Address
		Err       error
	})

	// Start workers.
	for chunkId := 0; chunkId < len(contractAddressesChunks); chunkId++ {
		// The chunks.
		contractAddressesChunk := contractAddressesChunks[chunkId]
		functionNamesChunk := functionNamesChunks[chunkId]
		functionArgsChunk := functionArgsChunks[chunkId]

		// Add one to the group.
		wg.Add(1)

		// Start routine.
		go func() {
			defer wg.Done()

			// Multi call factories..
			_, returnBytes, err := p.multiCall(abiList, contractAddressesChunk, functionNamesChunk, functionArgsChunk)
			if err != nil {
				resChan <- &struct {
					Reserves  [][]*big.Int
					Addresses []common.Address
					Err       error
				}{nil, nil, err}
				return
			}

			// Check errors.
			if len(returnBytes[0]) == 0 {
				resChan <- &struct {
					Reserves  [][]*big.Int
					Addresses []common.Address
					Err       error
				}{nil, nil, variables.EmptyResponse}
				return
			}

			// Decode pairs.
			var pairReserves = make([][]*big.Int, len(returnBytes))
			var pairAddrs = make([]common.Address, len(returnBytes))
			for k, addrBytes := range returnBytes {
				// Check if empty.

				for _, v := range addrBytes {
					if v != 0 && contractAddressesChunk[k].String() != "0x0000000000000000000000000000000000000000" {
						pairAddrs[k] = contractAddressesChunk[k]
						pairReserves[k] = make([]*big.Int, 2)
						pairReserves[k][0] = new(big.Int).SetBytes(addrBytes[0:32])
						pairReserves[k][1] = new(big.Int).SetBytes(addrBytes[32:64])

						// atomic.AddInt32(&shit, 1)
						// logger.Log.Infoln(fmt.Sprintf("[%d/%d] %s | %s, %s", atomic.LoadInt32(&shit), len(p.Pairs), pairAddrs[k].String(), pairReserves[k][0].String(), pairReserves[k][1].String()))
						break
					}
				}
			}

			resChan <- &struct {
				Reserves  [][]*big.Int
				Addresses []common.Address
				Err       error
			}{pairReserves, pairAddrs, nil}
		}()
	}

	// Wait.
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// Collect results.
	for r := range resChan {
		if r.Err != nil {
			return r.Err
		}
		for i, addr := range r.Addresses {
			if len(r.Reserves[i]) != 2 || r.Addresses[i].String() == "0x0000000000000000000000000000000000000000" {
				continue
			}
			p.AddressToPair[addr].SetReserves(r.Reserves[i][0], r.Reserves[i][1])
		}
	}

	return nil
}

// findDecimals
//	Finds the token decimals.
func (p *PairUpdater) findDecimals() error {
	// Address slice.
	var contractAddresses = make([]common.Address, len(p.params.Tokens.Addresses))
	for i, token := range p.params.Tokens.Addresses {
		contractAddresses[i] = token
	}

	// Function name slice.
	var functionNames = make([]string, len(p.params.Tokens.Addresses))
	for i, _ := range p.params.Tokens.Addresses {
		functionNames[i] = "decimals"
	}

	// Arguments slice.
	var functionArgs = make([][]interface{}, len(p.params.Tokens.Addresses))
	for i, _ := range p.params.Tokens.Addresses {
		functionArgs[i] = make([]interface{}, 0)
	}

	// Abi slice.
	var abiList = make([]abi.ABI, len(p.params.Tokens.Addresses))
	for i, _ := range p.params.Tokens.Addresses {
		abiList[i] = erc20Abi
	}

	// Multi call factories..
	_, returnBytes, err := p.multiCall(abiList, contractAddresses, functionNames, functionArgs)
	if err != nil {
		return err
	}

	// Check errors.
	if len(returnBytes) == 0 {
		return variables.EmptyResponse
	}

	p.TokenToDecimals = make(map[common.Address]*big.Int)

	// Decode pairs.
	for k, resBytes := range returnBytes {
		// Check if empty.
		for _, v := range resBytes {
			if v != 0 {
				addr := contractAddresses[k]
				p.TokenToDecimals[addr] = new(big.Int).SetBytes(resBytes)
				break
			}
		}
	}

	return nil
}

// findCircles
//	Finds DFS circles.
func (p *PairUpdater) findCircles() error {
	// Get token symbol.
	tempInSymbol, _ := p.params.Tokens.Symbols[p.params.Tokens.MainAddress]

	// DFS variables.
	path := make([]common.Address, 1)
	path[0] = p.params.Tokens.MainAddress

	pathSymbols := make([]string, 1)
	pathSymbols[0] = tempInSymbol

	dfsParams := DFSCircleParams{
		Path:          path,
		Symbols:       pathSymbols,
		Route:         make([]common.Address, 0),
		RouteFees:     make([]*big.Int, 0),
		RouteTokens:   make([][]common.Address, 0),
		RouteReserves: make([][]*big.Int, 0),
	}

	// Wait group.
	var wg sync.WaitGroup

	// The results map.
	var resCount int32 = 0
	var resMutex = sync.RWMutex{}

	p.Circles = make(map[uint64]*circle.Circle, 0)

	// Start DFS.
	wg.Add(1)
	dfsUtilOnlyCircle(dfsParams, &p.Circles, &resMutex, &resCount, &wg, p)
	wg.Wait()

	return nil
}

// sortCircles
//	Sorts DFS circles.
func (p *PairUpdater) sortCircles() {
	sortedTrades := p.quickSortCircles()
	p.sortedTrades.Swap(sortedTrades)
}

// subscribeToHeads
//	Subscribes to the new blocks.
func (p *PairUpdater) subscribeToHeads() {
	// Make new channel.
	p.blocksCh = make(chan *types.Header)

	// Subscribe to the new blocks.
	var err error
	p.blocksSub, err = variables.EthClient.SubscribeNewHead(context.Background(), p.blocksCh)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to subscribe to the new blocks.")
	}
}

// subscribeToPending
//	Subscribes to the new pending transactions.
func (p *PairUpdater) subscribeToPending() {
	// Make new channel.
	p.pendingCh = make(chan *common.Hash)

	// Subscribe to the new blocks.
	var err error
	p.pendingSub, err = variables.RpcClient.EthSubscribe(context.Background(), p.pendingCh, "newPendingTransactions")
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to subscribe to the pending transactions.")
	}
}

// listenBlocks
//	Listens new blocks.
func (p *PairUpdater) listenBlocks(header *types.Header) {
	// Total time to update a block.
	start := time.Now()

	// Get the previous 3 block's logs.
	prevLogs, err := p.backend.FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: p.PairAddresses,
		FromBlock: new(big.Int).SetUint64(p.lastBlockNum.Load().(uint64)),
		ToBlock:   header.Number,
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get previous logs.")
	}

	// Update events.
	for _, log := range prevLogs {
		// Iterate over topics.
		for _, topic := range log.Topics {
			// Continue if topic is not sync.
			if !bytes.Equal(syncId.Bytes(), topic.Bytes()) {
				continue
			}

			// Decode event.
			syncDetails, err := pairAbi.Unpack("Sync", log.Data)
			if err != nil {
				logger.Log.WithError(err).Errorln("Unable to decode 'sync' events.")
				break
			}

			// Update reserves.
			resA, resB := syncDetails[0].(*big.Int), syncDetails[1].(*big.Int)
			p.AddressToPair[log.Address].SetReserves(resA, resB)
			break
		}
	}

	// Sort the circles.
	p.sortCircles()

	// Update block number.
	p.lastBlockNum.Swap(header.Number.Uint64())

	// Clear the transaction history.
	p.TxHistoryMutex.Lock()
	p.PairToTxHistory = make(map[common.Address][]*types.Transaction)
	p.TxToOptionHistory = make(map[common.Hash]*circle.TradeOption)
	p.TxHistoryMutex.Unlock()

	// Set update time.
	updateTime := time.Since(start)

	// Print block latency.
	if variables.IsDev {
		go func() {
			logger.Log.Debugln("Block Latency:", updateTime, header.Number.Uint64())
		}()
	}

	// Check if it took too much time.
	if updateTime > 750*time.Millisecond {
		logger.Log.
			WithField("latency", time.Since(start)).
			WithField("blockNum", header.Number.Uint64()).
			Infoln("Block latency is too much! Skipping this block...")
	} else {
		// Trigger event.
		if p.OnSort != nil {
			p.OnSort(header, updateTime, p)
		}
	}
}

// listenPending gets triggered on a new pending transaction.
func (p *PairUpdater) listenPending(hash *common.Hash) {
	// Get transaction details.
	transaction, isPending, err := variables.EthClient.TransactionByHash(context.Background(), *hash)
	if err != nil || !isPending {
		return
	}

	// Filter out transactions.
	if transaction == nil || transaction.To() == nil || transaction.Data() == nil {
		return
	}
	if len(transaction.Data()) < 4 {
		return
	}
	if transaction.To().String() == "0x0000000000000000000000000000000000000000" ||
		transaction.To().String() == "0x000000000000000000000000000000000000dEaD" {
		return
	}
	if transaction.GasPrice().Cmp(variables.GasPrice) < 0 {
		return
	}

	// Skip if no history.
	if !variables.IsDev {
		p.TxHistoryMutex.RLock()
		if len(p.TxToOptionHistory) == 0 {
			p.TxHistoryMutex.RUnlock()
			return
		}
		p.TxHistoryMutex.RUnlock()
	}

	// Get from.
	msg, err := transaction.AsMessage(types.LatestSignerForChainID(variables.ChainId), big.NewInt(1))
	if err != nil {
		return
	}

	// Skip our transactions.
	if bytes.Equal(msg.From().Bytes(), variables.Wallet.Address().Bytes()) {
		return
	}

	// Marshall tx.
	simulationTx := make(map[string]string)
	simulationTx["from"] = msg.From().String()
	simulationTx["to"] = transaction.To().String()
	simulationTx["gas"] = fmt.Sprintf("0x%x", transaction.Gas())
	simulationTx["gasPrice"] = fmt.Sprintf("0x%x", transaction.GasPrice())
	simulationTx["value"] = fmt.Sprintf("0x%x", transaction.Value())
	simulationTx["data"] = "0x" + hex.EncodeToString(transaction.Data())
	simulationTx["nonce"] = fmt.Sprintf("0x%x", transaction.Nonce())

	// Options.
	options := make(map[string]interface{})
	options["disableStorage"] = true
	options["disableStack"] = false
	options["enableMemory"] = false
	options["timeout"] = "150ms"

	// Trace as call.
	var traceCallRes DebugTraceCall
	err = p.rpcBackend.Call(&traceCallRes, "debug_traceCall", simulationTx, "latest", options)
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("hash", transaction.Hash().String()).
			Errorln("Unable to simulate transaction.")
		return
	}

	// Iterate over trace call
	isSyncFound := false
	isPairFound := false
	var pairAddress common.Address
	for i, structLog := range traceCallRes.StructLogs {
		// Break if already found.
		if isSyncFound && isPairFound {
			break
		}

		// Check stack size.
		if i+2 > len(traceCallRes.StructLogs) || len(traceCallRes.StructLogs[i+1].Stack) < 3 {
			continue
		}

		// Catch events.
		if structLog.Op == "PUSH32" {
			// Get the pushed event id.
			eventId := traceCallRes.StructLogs[i+1].Stack[len(traceCallRes.StructLogs[i+1].Stack)-1]

			// Check if it's a "Sync" event.
			if strings.EqualFold(eventId, syncId.String()) {
				isSyncFound = true
				continue
			}

			// Check if it's a "Transfer" event.
			if strings.EqualFold(eventId, transferId.String()) {
				// Get the addresses.
				fromAddrRaw := traceCallRes.StructLogs[i+1].Stack[len(traceCallRes.StructLogs[i+1].Stack)-2]
				toAddrRaw := traceCallRes.StructLogs[i+1].Stack[len(traceCallRes.StructLogs[i+1].Stack)-3]

				// Skip if addresses are invalid.
				if !common.IsHexAddress(fromAddrRaw) || !common.IsHexAddress(toAddrRaw) {
					continue
				}

				// Check if any of the addresses are pair addresses.
				fromAddr, toAddr := common.HexToAddress(fromAddrRaw), common.HexToAddress(toAddrRaw)
				if slices.Contains(p.PairAddresses, fromAddr) {
					isPairFound = true
					pairAddress = fromAddr
				} else if slices.Contains(p.PairAddresses, toAddr) {
					isPairFound = true
					pairAddress = toAddr
				}
			}
		}
	}

	// Check if it's a swap that includes our pairs.
	if !isPairFound || !isSyncFound {
		return
	}

	// Lock the mutex.
	p.TxHistoryMutex.Lock()

	// Re-send the transaction.
	if tradeTxes, ok := p.PairToTxHistory[pairAddress]; ok {
		for i, tradeTx := range tradeTxes {
			// Continue if gas price is more.
			if tradeTx.GasPrice().Cmp(transaction.GasPrice()) > 0 {
				continue
			}

			// Get the options.
			tradeOption := p.TxToOptionHistory[tradeTx.Hash()]

			// Calculate the frontrun gas cost. (%15 more gas.)
			frontrunGasPrice := new(big.Int).Mul(transaction.GasPrice(), big.NewInt(115))
			frontrunGasPrice.Div(frontrunGasPrice, big.NewInt(100))

			// The frontrun gas cost. (gas price * gas)
			frontrunGasCost := new(big.Int).Mul(new(big.Int).SetUint64(transaction.Gas()), frontrunGasPrice)

			// Calculate the profit of the option.
			tradeProfit, err := tradeOption.LoanProfit()
			if err != nil {
				utils.PrintTradeOption(tradeOption)
				logger.Log.WithError(err).Errorln("Unable to calculate trade option's profit.")
				continue
			}

			// Log fields.
			logFields := logrus.Fields{
				"targetTx":        transaction.Hash(),
				"ourTx":           tradeTx.Hash(),
				"pairAddr":        pairAddress,
				"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToUnit(frontrunGasPrice, big.NewInt(9))),
			}

			// Check if we are still in profit.
			var replacedTx *types.Transaction
			if tradeProfit.Cmp(frontrunGasCost) > 0 {
				logger.Log.WithFields(logFields).Infoln("Updating trade transaction to frontrun competitors...")

				// Increase the gas price and resend transaction again.
				replacedTx = p.increaseTxGasPrice(tradeTx, frontrunGasPrice)

				// Replace.
				p.PairToTxHistory[pairAddress][i] = replacedTx
			} else {
				logger.Log.WithFields(logFields).Infoln("Trade transaction is not profitable anymore! Cancelling transaction...")

				// Calculate the cancel gas cost. (%15 more gas.)
				cancelGasPrice := new(big.Int).Mul(tradeTx.GasPrice(), big.NewInt(115))
				cancelGasPrice.Div(cancelGasPrice, big.NewInt(100))

				// Frontrun your own transaction and replace it with blank tx.
				p.cancelTx(tradeTx, cancelGasPrice)

				// Delete from history.
				p.PairToTxHistory[pairAddress] = append(p.PairToTxHistory[pairAddress][:i], p.PairToTxHistory[pairAddress][i+1:]...)
			}
		}
	}

	// Unlock the mutex.
	p.TxHistoryMutex.Unlock()
}

// increaseTxGasPrice
//	Increases the transaction's gas price and re-sends it again.
func (p *PairUpdater) increaseTxGasPrice(tx *types.Transaction, targetGasPrice *big.Int) *types.Transaction {
	// Replace transaction.
	replaceTransaction := types.NewTx(&types.LegacyTx{
		Nonce:    tx.Nonce(),
		To:       tx.To(),
		Value:    tx.Value(),
		Gas:      tx.Gas(),
		GasPrice: targetGasPrice,
		Data:     tx.Data(),
	})

	// New signer.
	signer, _ := variables.Wallet.NewTransactor()

	// Sign transaction.
	signedTx, err := signer.Signer(signer.From, replaceTransaction)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to sign replacement transaction.")
	}

	// Send the transaction.
	err = p.backend.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to send replacement transaction..")
	}

	logger.Log.WithFields(logrus.Fields{
		"replacedTx":      signedTx.Hash(),
		"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToUnit(targetGasPrice, big.NewInt(9))),
	}).Infoln("Replacement transaction sent!")

	return signedTx
	// TODO: broadcast message
}

// cancelTx
//	Increases the transaction's gas price and replaces it with blank transaction.
func (p *PairUpdater) cancelTx(tx *types.Transaction, targetGasPrice *big.Int) *types.Transaction {
	walletAddr := variables.Wallet.Address()

	// Blank transaction.
	blankTx := types.NewTx(&types.LegacyTx{
		Nonce:    tx.Nonce(),
		To:       &walletAddr,
		Value:    common.Big0,
		Gas:      21000,
		GasPrice: targetGasPrice,
		Data:     make([]byte, 0),
	})

	// New signer.
	signer, _ := variables.Wallet.NewTransactor()

	// Sign transaction.
	signedTx, err := signer.Signer(signer.From, blankTx)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to sign blank transaction.")
	}

	// Send the transaction.
	err = p.backend.SendTransaction(context.Background(), signedTx)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to send blank transaction..")
	}

	logger.Log.WithFields(logrus.Fields{
		"replacedTx":      signedTx.Hash(),
		"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToUnit(targetGasPrice, big.NewInt(9))),
	}).Infoln("Blank transaction sent!")

	// TODO: broadcast message
	return signedTx
}

// quickSortCircles
//	Quick sorts the circles.
func (p *PairUpdater) quickSortCircles() []*circle.TradeOption {
	// Get circles as array.
	tradeArr := make([]*circle.TradeOption, 0, len(p.Circles))
	for _, value := range p.Circles {
		// Calculate optimal in.
		optimalIn, amountsOut, err := p.GetOptimalIn(value)
		if err != nil && err != variables.NoArbitrage {
			logger.Log.WithError(err).Fatalln("Unable to sort trade options.")
		} else if err == variables.NoArbitrage {
			continue
		}

		option := &circle.TradeOption{
			Circle:     value,
			OptimalIn:  optimalIn,
			AmountsOut: amountsOut,
		}

		// Append to the list.
		tradeArr = append(tradeArr, option)
	}

	return p.quickSortUtil(tradeArr, 0, len(tradeArr)-1)
}

// quickSortUtil
//	Quick sort utility.
func (p *PairUpdater) quickSortUtil(arr []*circle.TradeOption, low, high int) []*circle.TradeOption {
	if low < high {
		var t int
		arr, t = p.partition(arr, low, high)
		arr = p.quickSortUtil(arr, low, t-1)
		arr = p.quickSortUtil(arr, t+1, high)
	}
	return arr
}

// partition
//	Quick sort utility.
func (p *PairUpdater) partition(arr []*circle.TradeOption, low, high int) ([]*circle.TradeOption, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		// Get profits.
		profitOne, err := arr[j].LoanProfit()
		if err != nil {
			utils.PrintTradeOption(arr[j])
			logger.Log.WithError(err).Fatalln("Unable to calculate trade profit.")
		}

		profitTwo, err := pivot.LoanProfit()
		if err != nil {
			utils.PrintTradeOption(pivot)
			logger.Log.WithError(err).Fatalln("Unable to calculate trade profit.")
		}

		if profitOne.Cmp(profitTwo) > 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

// dfsUtilDynamic
//	Helper function.
func dfsUtilOnlyCircle(params DFSCircleParams, circles *map[uint64]*circle.Circle, resMutex *sync.RWMutex, resCount *int32, wg *sync.WaitGroup, u *PairUpdater) {
	defer wg.Done()

	// Temporary tokens.
	tempIn := params.Path[len(params.Path)-1]

	// Get token pairs.
	tempInPairs, ok := u.TokenToPairs[tempIn]
	if !ok {
		panic("token pairs not found")
	}

	// Iterate over token pairs.
	for _, _pair := range tempInPairs {
		// Break if already found results.
		if len(params.Route) > config.Parsed.ArbitrageOptions.Limiters.MaxHops {
			break
		}

		// Skip if already visited.
		if slices.Contains(params.Route, _pair.Address) {
			continue
		}

		// Get pair tokens.
		pairTokens, ok := u.PairToTokens[_pair.Address]
		if !ok {
			panic("pair not in database")
		}

		// The output token.
		var tempOutToken common.Address

		// Find the output token.
		if bytes.Equal(pairTokens[0].Bytes(), tempIn.Bytes()) {
			tempOutToken = pairTokens[1]
		} else if bytes.Equal(pairTokens[1].Bytes(), tempIn.Bytes()) {
			tempOutToken = pairTokens[0]
		} else {
			panic("token addresses are not right for pair")
		}

		// Skip low liquidity.
		resIn, resOut, err := _pair.GetSortedReserves(tempOutToken)
		if err != nil || resIn.Cmp(big.NewInt(1e15)) <= 0 || resOut.Cmp(big.NewInt(1e15)) <= 0 {
			continue
		}

		// Get token symbol.
		tempOutSymbol, ok := u.params.Tokens.Symbols[tempOutToken]
		if !ok {
			panic("token symbol not found")
		}

		// Get route fee.
		routeFee, err := u.GetPairFee(_pair.Address)
		if err != nil {
			panic("pair fee not found")
		}

		// Check cycle.
		if len(params.Route) >= config.Parsed.ArbitrageOptions.Limiters.MinHops &&
			bytes.Equal(tempOutToken.Bytes(), params.Path[0].Bytes()) {
			// New variables.
			newPath := make([]common.Address, len(params.Path))
			newPathSymbols := make([]string, len(params.Symbols))
			newRoute := make([]common.Address, len(params.Route))
			newRouteFees := make([]*big.Int, len(params.Route))
			newRouteTokens := make([][]common.Address, len(params.Route))
			newRouteReserves := make([][]*big.Int, len(params.Route))

			// Copy old variables.
			copy(newPath, params.Path)
			copy(newPathSymbols, params.Symbols)
			copy(newRoute, params.Route)
			copy(newRouteFees, params.RouteFees)
			copy(newRouteTokens, params.RouteTokens)
			copy(newRouteReserves, params.RouteReserves)

			// Append new variables.
			newPath = append(newPath, tempOutToken)
			newPathSymbols = append(newPathSymbols, tempOutSymbol)
			newRouteFees = append(newRouteFees, routeFee)
			newRoute = append(newRoute, _pair.Address)
			newRouteTokens = append(newRouteTokens, pairTokens)
			newRouteReserves = append(newRouteReserves, []*big.Int{_pair.ReserveA, _pair.ReserveB})

			// Get route as structs.
			tmpPairs := make([]*dexpair.DexPair, len(newRoute))
			for m, pairAddr := range newRoute {
				tmpPairs[m] = u.AddressToPair[pairAddr]
			}

			// New circle.
			arbCircle := &circle.Circle{
				Path:          newPath,
				Symbols:       newPathSymbols,
				Pairs:         tmpPairs,
				PairAddresses: newRoute,
				PairFees:      newRouteFees,
				PairTokens:    newRouteTokens,
				PairReserves:  newRouteReserves,
			}

			// Check limiter.
			resMutex.RLock()
			if *resCount > config.Parsed.ArbitrageOptions.Limiters.MaxCircles {
				resMutex.RUnlock()
				return
			}
			resMutex.RUnlock()

			// Increase counter.
			resMutex.Lock()
			circleId := arbCircle.ID()
			(*circles)[circleId] = arbCircle
			*resCount += 1
			resMutex.Unlock()
		} else {
			// The params.
			newParams := DFSCircleParams{
				Path:          make([]common.Address, len(params.Path)),
				Symbols:       make([]string, len(params.Symbols)),
				Route:         make([]common.Address, len(params.Route)),
				RouteFees:     make([]*big.Int, len(params.Route)),
				RouteTokens:   make([][]common.Address, len(params.Route)),
				RouteReserves: make([][]*big.Int, len(params.Route)),
			}
			copy(newParams.Path, params.Path)
			copy(newParams.Symbols, params.Symbols)
			copy(newParams.Route, params.Route)
			copy(newParams.RouteFees, params.RouteFees)
			copy(newParams.RouteTokens, params.RouteTokens)
			copy(newParams.RouteReserves, params.RouteReserves)

			newParams.Path = append(newParams.Path, tempOutToken)
			newParams.Symbols = append(newParams.Symbols, tempOutSymbol)
			newParams.Route = append(params.Route, _pair.Address)
			newParams.RouteFees = append(params.RouteFees, routeFee)
			newParams.RouteTokens = append(params.RouteTokens, pairTokens)
			newParams.RouteReserves = append(newParams.RouteReserves, []*big.Int{_pair.ReserveA, _pair.ReserveB})

			// Recursive
			wg.Add(1)
			go dfsUtilOnlyCircle(newParams, circles, resMutex, resCount, wg, u)
		}
	}
}

// multiCall
// 	Calls the multiple same contracts and returns the responses. (Max 21,000 gas limit.)
func (p *PairUpdater) multiCall(
	contractAbis []abi.ABI,
	contractAddresses []common.Address,
	functionNames []string,
	functionArgs [][]interface{},
) (*big.Int, [][]byte, error) {
	// Iterate through the addresses.
	var calls []abis.MulticallCall
	for i, contractAddr := range contractAddresses {
		// Create new empty byte array.
		inputBytes, err := contractAbis[i].Pack(functionNames[i], functionArgs[i]...)
		if err != nil {
			return nil, nil, err
		}

		// Ready the call.
		call := abis.MulticallCall{
			Target:   contractAddr,
			CallData: inputBytes,
		}

		calls = append(calls, call)
	}

	// Get the call bytes.
	callBytes, err := multicallerAbi.Pack("aggregate", calls)
	if err != nil {
		return nil, nil, err
	}

	// Create new blank message.
	msg := ethereum.CallMsg{
		To:   &p.params.Multicaller.Address,
		Data: callBytes,
	}

	// Call the aggregate function.
	gasLimit, err := p.backend.EstimateGas(context.Background(), msg)
	result, err := p.backend.PendingCallContract(context.Background(), msg)
	if err != nil {
		return nil, nil, err
	}

	_ = gasLimit

	// Get the results.
	out, err := multicallerAbi.Unpack("aggregate", result)
	if err != nil {
		return nil, nil, err
	}

	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), out[1].([][]byte), nil
}

// chunkBy
// 	Splits items into chunks.
func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	// While there are more items remaining than chunkSize...
	for chunkSize < len(items) {
		// We take a slice of size chunkSize from the items array and append it to the new array
		chunks = append(chunks, items[0:chunkSize])
		// Then we remove those elements from the items array
		items = items[chunkSize:]
	}
	// Finally, we append the remaining items to the new array and return it
	return append(chunks, items)
}
