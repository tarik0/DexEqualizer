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
	"github.com/psilva261/timsort/v2"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/circle"
	config "github.com/tarik0/DexEqualizer/config"
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

	// Multi call factories.
	callBlockNum := new(big.Int).SetUint64(p.highestBlockNum.Load().(uint64))
	_, returnBytes, err := p.multiCall(abiList, contractAddresses, functionNames, functionArgs, callBlockNum)
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
		callBlockNum := new(big.Int).SetUint64(p.highestBlockNum.Load().(uint64))
		go func(addresses []common.Address, funcNames []string, funcArgs [][]interface{}) {
			defer wg.Done()

			// Multi call factories..
			_, returnBytes, err := p.multiCall(abiList, addresses, funcNames, funcArgs, callBlockNum)
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
		contractAddresses[i] = _pair.Address()
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

	// Get current block number.
	callBlockNum := new(big.Int).SetUint64(p.highestBlockNum.Load().(uint64))

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

			// Multi call factories.
			_, returnBytes, err := p.multiCall(abiList, contractAddressesChunk, functionNamesChunk, functionArgsChunk, callBlockNum)
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
			p.AddressToPair[addr].SetReserves(r.Reserves[i][0], r.Reserves[i][1], callBlockNum)
		}
	}

	logger.Log.WithField("blockNum", callBlockNum).Debugln("Pool reserves are found!")

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

	// Multi call factories.
	callBlockNum := new(big.Int).SetUint64(p.highestBlockNum.Load().(uint64))
	_, returnBytes, err := p.multiCall(abiList, contractAddresses, functionNames, functionArgs, callBlockNum)
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

	// The results and the wait group.
	var circleCh = make(chan *circle.Circle, config.Parsed.ArbitrageOptions.Limiters.MaxCircles)
	var mutex = sync.RWMutex{}
	var wg = sync.WaitGroup{}

	// Start DFS.
	wg.Add(1)
	dfsUtilOnlyCircle(dfsParams, circleCh, &mutex, &wg, p)
	wg.Wait()
	close(circleCh)

	// Limit circle amount.
	circleLen := len(circleCh)
	if circleLen > int(config.Parsed.ArbitrageOptions.Limiters.MaxCircles) {
		circleLen = int(config.Parsed.ArbitrageOptions.Limiters.MaxCircles)
	}
	p.Circles = make(map[uint64]*circle.Circle, circleLen)

	// Get the results.
	for _circle := range circleCh {
		_id := _circle.ID()
		if _, ok := p.Circles[_id]; !ok {
			p.Circles[_id] = _circle
			logger.Log.Debugln(len(p.Circles), _circle.SymbolsStr())
			logger.Log.Debugln(_circle.PairAddressesStr())
		}
	}

	return nil
}

// sortCircles
//	Sorts DFS circles.
func (p *PairUpdater) sortCircles() []*circle.TradeOption {
	sortedTrades := p.quickSortCircles()
	return sortedTrades
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
	// Test if `debug_traceCall` is supported.
	supportedModules, err := p.rpcBackend.SupportedModules()
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to check supported modules.")
		return
	}
	if _, ok := supportedModules["debug"]; !ok {
		logger.Log.Warningln("'debug_traceCall' is not supported. Transaction simulator disabled!")
		return
	}

	// Make new channel.
	p.pendingCh = make(chan *common.Hash)

	// Subscribe to the new blocks.
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

	// Get the latest block.
	lastSyncBlockNum := p.GetLastSyncBlockNumber()

	// From block.
	fromBlock := new(big.Int).Set(header.Number)
	if header.Number.Uint64() != lastSyncBlockNum {
		fromBlock.SetUint64(lastSyncBlockNum)
	}

	// Get the previous block logs.
	prevLogs, err := p.backend.FilterLogs(context.Background(), ethereum.FilterQuery{
		Addresses: p.PairAddresses,
		FromBlock: fromBlock,
		ToBlock:   header.Number,
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get previous logs.")
	}

	// Update reserves.
	p.filterLogsMutex.Lock()
	for _, log := range prevLogs {
		// Continue if not in addresses.
		if _, ok := p.AddressToPair[log.Address]; !ok {
			continue
		}

		// Iterate over topics.
		for _, topic := range log.Topics {
			// Continue if topic is not sync.
			if !bytes.EqualFold(syncId.Bytes(), topic.Bytes()) {
				continue
			}

			// Decode event.
			syncDetails, err := pairAbi.Unpack("Sync", log.Data)
			if err != nil {
				logger.Log.WithError(err).Fatalln("Unable to decode 'sync' events.")
				break
			}

			// Update reserves.
			resA, resB := syncDetails[0].(*big.Int), syncDetails[1].(*big.Int)
			p.AddressToPair[log.Address].SetReserves(resA, resB, new(big.Int).SetUint64(log.BlockNumber))
			break
		}
	}
	p.lastSyncBlockNum.Swap(header.Number.Uint64())
	p.filterLogsMutex.Unlock()

	updateTime := time.Since(start)
	logger.Log.
		WithField("updateTime", updateTime).
		WithField("fromBlock", fromBlock).
		WithField("toBlock", header.Number).
		Debugln("Synced with the block.")

	// Sort the circles.
	sortStart := time.Now()
	sortedTradeOptions := p.sortCircles()
	sortTime := time.Since(sortStart)

	// Trigger sort listener.
	p.listenSort(sortedTradeOptions, header, updateTime, sortTime)
}

// listenSort
//	Listens new block updates.
func (p *PairUpdater) listenSort(options []*circle.TradeOption, header *types.Header, updateTime time.Duration, sortTime time.Duration) {
	// Check if block has already passed.
	highestBlockNum := p.GetHighestBlockNumber()
	if highestBlockNum > header.Number.Uint64() || updateTime+sortTime > 2*time.Second {
		logger.Log.
			WithField("updateTime", updateTime).
			WithField("blockNum", header.Number.Uint64()).
			WithField("latestBlockNum", highestBlockNum).
			Debugln("Block latency is too much! Skipping this block...")
		return
	}

	// Skip if no trades.
	if len(options) == 0 {
		return
	}

	// Limit options to best 5.
	if len(options) > 5 {
		options = options[:5]
	}

	// Trigger event.
	if p.OnSort != nil && len(options) > 0 {
		p.OnSort(header, options, sortTime+updateTime, p)
	}
}

// listenPending gets triggered on a new pending transaction.
func (p *PairUpdater) listenPending(hash *common.Hash) {
	// Skip if no hash.
	if hash == nil {
		return
	}

	// Get transaction details.
	transaction, isPending, err := variables.EthClient.TransactionByHash(context.Background(), *hash)
	if err != nil || transaction == nil || isPending != false {
		return
	}

	// Get from.
	msg, err := transaction.AsMessage(types.LatestSignerForChainID(variables.ChainId), big.NewInt(1))
	if err != nil {
		return
	}

	// Skip our transactions.
	if bytes.EqualFold(msg.From().Bytes(), variables.Wallet.Address().Bytes()) {
		return
	}

	// Check account nonce.
	val, ok := p.accountToPendingTx.Load(msg.From())

	// Check if there are already a pending transaction for that account.
	if ok {
		// Previous account transaction.
		prevAccountTx := val.(*types.Transaction)

		// Check nonce's.
		if transaction.Nonce() == prevAccountTx.Nonce() {
			// Check gas prices.
			if transaction.GasPrice().Cmp(prevAccountTx.GasPrice()) > 0 {
				// New transaction has same nonce but more gas.
				// Replace the transaction.
				p.accountToPendingTx.Store(msg.From(), transaction)
			} else {
				// New transaction has same nonce but less or same gas.
				// Skip.
				return
			}
		} else if transaction.Nonce() > prevAccountTx.Nonce() {
			// Replace the transaction.
			p.accountToPendingTx.Store(msg.From(), transaction)
		} else {
			// Skip if nonce is less.
			return
		}
	} else {
		// Add it to the transactions.
		p.accountToPendingTx.Store(msg.From(), transaction)
	}

	// Filter out transactions.
	if transaction.To() == nil || transaction.Data() == nil {
		return
	}
	if len(transaction.Data()) < 4 {
		return
	}
	if transaction.To().String() == "0x0000000000000000000000000000000000000000" ||
		transaction.To().String() == "0x000000000000000000000000000000000000dEaD" {
		return
	}
	if transaction.GasPrice().Cmp(variables.GasPrice) < 0 || transaction.Gas() < 70000 {
		return
	}

	// Static search.
	if !p.staticSearch(transaction) {
		// Dynamic search.
		p.dynamicSearch(msg, transaction)
	}
}

// staticSearch
//	Searches the transaction data and check if any of our pairs are used.
func (p *PairUpdater) staticSearch(transaction *types.Transaction) bool {
	// Limit call data to 100 + 4 bytes
	callData := transaction.Data()
	if len(callData) > 104 {
		callData = callData[4:100]
	}

	// Iterate over pairs.
	isFound := false
	for _, pairAddr := range p.PairAddresses {
		// Trigger search if pair address is in the transaction data.
		if bytes.Contains(callData, pairAddr.Bytes()) {
			isFound = true
			p.TxHistorySearch <- struct {
				TargetTx       *types.Transaction
				TargetPairAddr common.Address
			}{
				TargetTx:       transaction,
				TargetPairAddr: pairAddr,
			}
		}
	}

	return isFound
}

// dynamicSearch
// 	Simulates a transaction with `debug_traceCall` and checks if any of our pairs are used.
func (p *PairUpdater) dynamicSearch(msg types.Message, transaction *types.Transaction) {
	// Marshall tx.
	simulationTx := make(map[string]string)
	simulationTx["from"] = msg.From().String()
	simulationTx["to"] = transaction.To().String()
	simulationTx["gas"] = fmt.Sprintf("0x%x", transaction.Gas())
	simulationTx["gasPrice"] = fmt.Sprintf("0x%x", transaction.GasPrice())
	simulationTx["value"] = fmt.Sprintf("0x%x", transaction.Value())
	simulationTx["data"] = "0x" + hex.EncodeToString(transaction.Data())
	simulationTx["nonce"] = fmt.Sprintf("0x%x", transaction.Nonce())

	// Tracer options.
	tracerOptions := make(map[string]interface{})
	tracerOptions["onlyTopCall"] = "false"
	tracerOptions["withLog"] = "true"

	// Options.
	options := make(map[string]interface{})
	options["disableStorage"] = true
	options["disableStack"] = false
	options["enableMemory"] = false
	options["timeout"] = "75ms"
	options["tracer"] = "callTracer"
	options["tracerConfig"] = tracerOptions

	// Trace as call.
	var traceCallRes DebugTraceCall
	err := p.rpcBackend.Call(&traceCallRes, "debug_traceCall", simulationTx, "latest", options)
	if err != nil {
		logger.Log.
			WithError(err).
			WithField("hash", transaction.Hash().String()).
			Debugln("Unable to simulate transaction.")
		return
	}

	// Pair addresses channel.
	var wg sync.WaitGroup
	var pairAddrsCh chan common.Address

	// Recursive checker function.
	var checkCall func(DebugTraceCall)
	checkCall = func(call DebugTraceCall) {
		defer wg.Done()
		if slices.Contains(p.PairAddresses, call.To) {
			pairAddrsCh <- call.To
		}

		// Check sub calls.
		if call.Calls != nil && len(call.Calls) > 0 {
			for _, subCall := range call.Calls {
				wg.Add(1)
				go checkCall(subCall)
			}
		}
	}

	// Wait checker.
	wg.Add(1)
	go checkCall(traceCallRes)
	wg.Wait()

	// Send to the search channel.
	for pairAddr := range pairAddrsCh {
		p.TxHistorySearch <- struct {
			TargetTx       *types.Transaction
			TargetPairAddr common.Address
		}{
			TargetTx:       transaction,
			TargetPairAddr: pairAddr,
		}
	}
}

// listenHistory gets triggered when a transaction gets added to history channel.
func (p *PairUpdater) listenHistory() {
	// Generate new channels.
	p.TxHistoryReset = make(chan bool)
	p.TxHistoryAdd = make(chan struct {
		Tx          *types.Transaction
		Option      *circle.TradeOption
		BlockNumber *big.Int
	})
	p.TxHistorySearch = make(chan struct {
		TargetTx       *types.Transaction
		TargetPairAddr common.Address
	})

	// Generate new history.
	p.hashToOptionHistory = make(map[common.Hash]*circle.TradeOption)
	p.hashToTxHistory = make(map[common.Hash]*types.Transaction)
	p.hashToTxBlock = make(map[common.Hash]*big.Int)

	go func() {
		for {
			select {
			// Reset channel is prioritized.
			case _ = <-p.TxHistoryReset:
				// Clear account pending history.
				p.accountToPendingTx.Range(func(key interface{}, value interface{}) bool {
					p.accountToPendingTx.Delete(key)
					return true
				})

				// Reset the history.
				if len(p.hashToTxHistory) == 0 {
					continue
				}

				logger.Log.WithField("historyLen", len(p.hashToTxHistory)).Debugln("History cleared!")
				p.hashToOptionHistory = make(map[common.Hash]*circle.TradeOption)
				p.hashToTxHistory = make(map[common.Hash]*types.Transaction)
				p.hashToTxBlock = make(map[common.Hash]*big.Int)
			default:
			case txInfo := <-p.TxHistoryAdd:
				// Iterate over pair addresses.
				p.hashToOptionHistory[txInfo.Tx.Hash()] = txInfo.Option
				p.hashToTxHistory[txInfo.Tx.Hash()] = txInfo.Tx
				p.hashToTxBlock[txInfo.Tx.Hash()] = txInfo.BlockNumber
				logger.Log.WithField("historyLen", len(p.hashToTxHistory)).Debugln("Added to the history!")
			case searchInfo := <-p.TxHistorySearch:
				// Skip if history is empty.
				if len(p.hashToOptionHistory) == 0 {
					continue
				}

				// Search history.
				for prevTxHash, prevOption := range p.hashToOptionHistory {
					// Get previous transaction's block number.
					prevTxBlock := p.hashToTxBlock[prevTxHash]
					latestBlock := new(big.Int).SetUint64(p.highestBlockNum.Load().(uint64))
					if latestBlock.Cmp(prevTxBlock) != 0 {
						break
					}

					// Continue if none of the pairs are used in that transaction.
					if !slices.Contains(prevOption.Circle.PairAddresses, searchInfo.TargetPairAddr) {
						continue
					}

					// Get the previous transaction.
					prevTx := p.hashToTxHistory[prevTxHash]
					tradeOption := p.hashToOptionHistory[prevTxHash]

					// Continue if block has already passed.
					if prevTxBlock.Uint64() >= p.GetHighestBlockNumber() {
						logger.Log.WithField("hash", prevTxHash.String()).Infoln("Block has already passed! Skipping this transaction...")
						continue
					}

					// Continue if gas price is lower.
					if prevTx.GasPrice().Cmp(searchInfo.TargetTx.GasPrice()) > 0 {
						continue
					}

					// Calculate the frontrun gas cost. (%15 more gas.)
					frontrunGasPrice := new(big.Int).Mul(searchInfo.TargetTx.GasPrice(), big.NewInt(115))
					frontrunGasPrice.Div(frontrunGasPrice, big.NewInt(100))

					// Calculate the profit limit of the option.
					newTradeProfitLimit := prevOption.GetTradeCost(frontrunGasPrice)
					tradeProfit, err := prevOption.NormalProfit()
					if err != nil {
						logger.Log.WithError(err).Errorln("Unable to calculate trade profit.")
						utils.PrintTradeOption(prevOption)
						continue
					}

					// Get from.
					prevAccountMsg, err := searchInfo.TargetTx.AsMessage(types.LatestSignerForChainID(variables.ChainId), big.NewInt(1))
					if err != nil {
						logger.Log.WithError(err).Fatalln("Unable to get transaction as message.")
					}

					// Log fields.
					logFields := logrus.Fields{
						"targetTx":        searchInfo.TargetTx.Hash(),
						"account":         prevAccountMsg.From(),
						"nonce":           prevAccountMsg.Nonce(),
						"ourTx":           prevTx.Hash(),
						"pairAddr":        searchInfo.TargetPairAddr,
						"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToUnit(frontrunGasPrice, big.NewInt(9))),
					}

					// Check if we are still in profit.
					if tradeProfit.Cmp(newTradeProfitLimit) >= 0 {
						logger.Log.WithFields(logFields).Infoln("Updating trade transaction to frontrun competitors...")
						logger.Log.Infoln("")

						// Increase the gas price and resend transaction again.
						replacedTx := p.increaseTxGasPrice(prevTx, tradeOption, prevTxBlock, frontrunGasPrice)

						// Replace.
						if replacedTx != nil {
							delete(p.hashToOptionHistory, prevTxHash)
							delete(p.hashToTxHistory, prevTxHash)
							delete(p.hashToTxBlock, prevTxHash)
							p.hashToOptionHistory[replacedTx.Hash()] = prevOption
							p.hashToTxHistory[replacedTx.Hash()] = replacedTx
							p.hashToTxBlock[replacedTx.Hash()] = prevTxBlock
						}
					} else {
						// Calculate the cancel gas cost. (%15 more gas.)
						cancelGasPrice := new(big.Int).Mul(prevTx.GasPrice(), big.NewInt(115))
						cancelGasPrice.Div(cancelGasPrice, big.NewInt(100))

						// Enable auto-cancel
						if cancelGasPrice.Cmp(variables.CancelThresholdGasPrice) > 0 {
							logger.Log.WithFields(logFields).Infoln("Trade transaction might not be profitable anymore! Cancelling transaction...")

							// Frontrun your own transaction and replace it with blank tx.
							replacedTx := p.cancelTx(prevTx, tradeOption, prevTxBlock, cancelGasPrice)

							// Delete from history.
							if replacedTx != nil {
								delete(p.hashToOptionHistory, prevTx.Hash())
								delete(p.hashToTxHistory, prevTxHash)
								delete(p.hashToTxBlock, prevTxHash)
							}
						}
					}
				}
			}
		}
	}()
}

// increaseTxGasPrice
//	Increases the transaction's gas price and re-sends it again.
func (p *PairUpdater) increaseTxGasPrice(tx *types.Transaction, option *circle.TradeOption, prevBlock *big.Int, targetGasPrice *big.Int) *types.Transaction {

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
		"oldTx":           tx.Hash(),
		"newTx":           signedTx.Hash(),
		"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToUnit(targetGasPrice, big.NewInt(9))),
	}).Infoln("Replacement transaction sent!")

	// Broadcast message.
	err = variables.Hub.BroadcastMsg(
		fmt.Sprintf("Transaction's gas price got updated! (%.3f Gwei)", utils.WeiToUnit(targetGasPrice, big.NewInt(9))),
	)
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to broadcast message.")
	}

	return signedTx
}

// cancelTx
//	Increases the transaction's gas price and replaces it with blank transaction.
func (p *PairUpdater) cancelTx(tx *types.Transaction, option *circle.TradeOption, prevBlock *big.Int, targetGasPrice *big.Int) *types.Transaction {
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
		logger.Log.WithError(err).Errorln("Unable to send blank transaction..")
		return nil
	}

	logger.Log.WithFields(logrus.Fields{
		"replacedTx":      signedTx.Hash(),
		"updatedGasPrice": fmt.Sprintf("%.3f Gwei", utils.WeiToUnit(targetGasPrice, big.NewInt(9))),
	}).Infoln("Blank transaction sent!")

	// Broadcast message.
	err = variables.Hub.BroadcastMsg("Transaction got canceled because It's not profitable anymore!")
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to broadcast message.")
	}

	return signedTx
}

// quickSortCircles
//	Quick sorts the circles.
func (p *PairUpdater) quickSortCircles() []*circle.TradeOption {
	// Get circles as array.
	tradeArr := make([]interface{}, 0)
	for _, value := range p.Circles {
		// Calculate optimal in.
		optimalIn, amountsOut, err := p.GetOptimalIn(value)
		if err != nil && err != variables.NoArbitrage {
			logger.Log.WithError(err).Fatalln("Unable to sort trade options.")
		} else if err == variables.NoArbitrage {
			continue
		}

		// Generate new trade option.
		option, err := circle.NewTradeOption(value, optimalIn, amountsOut)
		if err != nil {
			continue
		}

		// Append to the list.
		tradeArr = append(tradeArr, option)
	}

	// Sort with timsort.
	timsort.Sort(tradeArr, func(a, b interface{}) bool {
		// Get profits.
		profitOne, err := a.(*circle.TradeOption).NormalProfit()
		if err != nil {
			utils.PrintTradeOption(a.(*circle.TradeOption))
			logger.Log.WithError(err).Fatalln("Unable to calculate trade profit.")
		}

		profitTwo, err := b.(*circle.TradeOption).NormalProfit()
		if err != nil {
			utils.PrintTradeOption(b.(*circle.TradeOption))
			logger.Log.WithError(err).Fatalln("Unable to calculate trade profit.")
		}

		return profitOne.Cmp(profitTwo) > 0
	})

	// Convert interface to trade.
	tmp := make([]*circle.TradeOption, len(tradeArr))
	for i, v := range tradeArr {
		tmp[i] = v.(*circle.TradeOption)
	}
	return tmp

	// The quicksort.
	// return p.quickSortUtil(tradeArr, 0, len(tradeArr)-1)
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
		profitOne, err := arr[j].NormalProfit()
		if err != nil {
			utils.PrintTradeOption(arr[j])
			logger.Log.WithError(err).Fatalln("Unable to calculate trade profit.")
		}

		profitTwo, err := pivot.NormalProfit()
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

// dfsUtilOnlyCircle
//	Helper function.
func dfsUtilOnlyCircle(params DFSCircleParams, resultsCh chan *circle.Circle, mutex *sync.RWMutex, wg *sync.WaitGroup, u *PairUpdater) {
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

		// Break if already results found.
		mutex.RLock()
		if len(resultsCh) == cap(resultsCh) {
			mutex.RUnlock()
			break
		}
		mutex.RUnlock()

		// Skip if already visited.
		if slices.Contains(params.Route, _pair.Address()) {
			continue
		}

		// Get pair tokens.
		pairTokens, ok := u.PairToTokens[_pair.Address()]
		if !ok {
			panic("pair not in database")
		}

		// The output token.
		var tempOutToken common.Address

		// Find the output token.
		if bytes.EqualFold(pairTokens[0].Bytes(), tempIn.Bytes()) {
			tempOutToken = pairTokens[1]
		} else if bytes.EqualFold(pairTokens[1].Bytes(), tempIn.Bytes()) {
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
		routeFee, err := u.GetPairFee(_pair.Address())
		if err != nil {
			panic("pair fee not found")
		}

		// Check cycle.
		if len(params.Route) >= config.Parsed.ArbitrageOptions.Limiters.MinHops &&
			bytes.EqualFold(tempOutToken.Bytes(), params.Path[0].Bytes()) {
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
			newRoute = append(newRoute, _pair.Address())
			newRouteTokens = append(newRouteTokens, pairTokens)
			newRouteReserves = append(newRouteReserves, _pair.GetReserves())

			// Get route as structs.
			tmpPairs := make([]*dexpair.DexPair, len(newRoute))
			for m, pairAddr := range newRoute {
				tmpPairs[m] = u.AddressToPair[pairAddr]
			}

			// New circle.
			arbCircle, err := circle.NewCircle(
				newPath,
				newPathSymbols,
				tmpPairs,
				newRouteFees,
				newRouteTokens,
				newRoute,
				newRouteReserves,
			)
			if err != nil {
				logger.Log.WithError(err).Fatalln("Unable to generate new circle.")
			}

			// Break if already results found.
			mutex.Lock()
			if len(resultsCh) == cap(resultsCh) {
				mutex.Unlock()
				break
			}

			// Send circle to the channel.
			resultsCh <- arbCircle
			mutex.Unlock()
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
			newParams.Route = append(params.Route, _pair.Address())
			newParams.RouteFees = append(params.RouteFees, routeFee)
			newParams.RouteTokens = append(params.RouteTokens, pairTokens)
			newParams.RouteReserves = append(newParams.RouteReserves, _pair.GetReserves())

			// Recursive
			wg.Add(1)
			go dfsUtilOnlyCircle(newParams, resultsCh, mutex, wg, u)
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
	blockNumber *big.Int,
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
	result, err := p.backend.CallContract(context.Background(), msg, blockNumber)
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
