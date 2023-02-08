package updater

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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
	_, returnBytes, err := p.multiCallBatch(abiList, contractAddresses, functionNames, functionArgs, new(big.Int).Set(common.Big0))
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

	// Set arrays and maps.
	p.TokenToPairs = make(map[common.Address][]*dexpair.DexPair)
	p.AddressToPair = make(map[common.Address]*dexpair.DexPair)
	p.PairToTokens = make(map[common.Address][]common.Address)
	p.PairToFactory = make(map[common.Address]common.Address)
	p.Pairs = make([]*dexpair.DexPair, 0)

	// Batch multicall.
	_, callResponses, err := p.multiCallBatch(abiList, contractAddresses, functionNames, functionArgs, new(big.Int).Set(common.Big0))
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to find pair addresses.")
	}

	// Check response length.
	if len(callResponses) != totalPairsSize {
		panic("missing multicall response")
	}

	// Iterate over response.
	for i, returnBytes := range callResponses {
		// Decode pair address.
		pairAddr := common.BytesToAddress(returnBytes)

		// Skip dead address.
		if strings.EqualFold(pairAddr.String(), "0x0000000000000000000000000000000000000000") ||
			strings.EqualFold(pairAddr.String(), "0x000000000000000000000000000000000000dEaD") {
			continue
		}

		// The other arguments.
		tokenA := (functionArgs[i][0]).(common.Address)
		tokenB := (functionArgs[i][1]).(common.Address)
		factory := contractAddresses[i]

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
		p.PairToTokens[pairAddr] = dexpair.SortTokens(tokenA, tokenB)
		p.PairToFactory[pairAddr] = factory
		p.Pairs = append(p.Pairs, tmpPair)
	}

	return nil
}

// findReserves
//	Finds the pair reserves from the pairs.
func (p *PairUpdater) findReserves(blockNum *big.Int) (*big.Int, error) {
	// Skip if no pair.
	if len(p.Pairs) == 0 {
		return nil, variables.NoPairFound
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

	// Batch multicall.
	resBlockNum, callResponses, err := p.multiCallBatch(abiList, p.PairAddresses, functionNames, functionArgs, blockNum)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to find pair reserves.")
	}

	// Iterate over response.
	for i, returnBytes := range callResponses {
		pairAddr := p.PairAddresses[i]

		// Continue if empty.
		if len(returnBytes) != 96 {
			logger.Log.
				WithField("returnBytesLen", len(returnBytes)).
				WithField("pair", pairAddr.String()).
				Errorln("getReserve response is not right for this pair.")
			continue
		}

		// Decode reserves.
		reserve0 := new(big.Int).SetBytes(returnBytes[0:32])
		reserve1 := new(big.Int).SetBytes(returnBytes[32:64])

		// Check reserves.
		if (reserve0.Cmp(common.Big0) == 0 && reserve1.Cmp(common.Big0) != 0) ||
			(reserve1.Cmp(common.Big0) == 0 && reserve0.Cmp(common.Big0) != 0) {
			panic("invalid reserves")
		}

		// Send to channel.
		p.syncCh <- SyncAction{
			Address:     pairAddr,
			Res0:        reserve0,
			Res1:        reserve1,
			BlockNumber: resBlockNum,
			TxIndex:     abi.MaxUint256,
			LogIndex:    abi.MaxUint256,
		}
	}

	return resBlockNum, nil
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
	_, returnBytes, err := p.multiCallBatch(abiList, contractAddresses, functionNames, functionArgs, new(big.Int).Set(common.Big0))
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
		Path:        path,
		Symbols:     pathSymbols,
		Route:       make([]common.Address, 0),
		RouteFees:   make([]*big.Int, 0),
		RouteTokens: make([][]common.Address, 0),
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

// subscribeToLogs
//	Subscribes to new logs.
func (p *PairUpdater) subscribeToLogs(fromBlock uint64) {
	// Make new channel.
	p.logsCh = make(chan types.Log)

	// Subscribe to the new blocks.
	var err error
	p.logsSub, err = variables.EthClient.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock - 1),
		Addresses: p.PairAddresses,
	}, p.logsCh)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to subscribe to the new blocks.")
	}
}

// subscribeToPending
//	Subscribes to the new pending transactions.
func (p *PairUpdater) subscribeToPending() {
	// Test if `debug_traceCall` is supported.
	supportedModules, err := p.pendingBackend.SupportedModules()
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

// compareAndSwapGasPrice
//	Compares the previous min. gas price for the pair.
func (p *PairUpdater) compareAndSwapGasPrice(pairAddr common.Address, tx *types.Transaction) {
	// Lock the gas price mutex.
	p.pairToMinGasPriceMutex.Lock()

	// Update the min. gas price for the pair.
	minGasPrice, ok := p.pairToMinGasPrice[pairAddr]

	// The log fields.
	logField := logrus.Fields{
		"pair": pairAddr,
	}

	// Check if pair already has a record.
	if ok && tx.GasPrice().Cmp(minGasPrice) > 0 {
		// Compare the gas price.
		logField["oldGasPrice"] = fmt.Sprintf("%.2f Gwei", utils.WeiToGwei(minGasPrice))
		logField["newGasPrice"] = fmt.Sprintf("%.2f Gwei", utils.WeiToGwei(tx.GasPrice()))
		p.pairToMinGasPrice[pairAddr] = tx.GasPrice()
		logger.Log.WithFields(logField).Debugln("Updated gas requirement for the pair.")
	} else if !ok {
		// Set the min gas price.
		p.pairToMinGasPrice[pairAddr] = tx.GasPrice()
		logField["newGasPrice"] = fmt.Sprintf("%.2f Gwei", utils.WeiToGwei(tx.GasPrice()))
		logger.Log.WithFields(logField).Debugln("Updated gas requirement for the pair.")
	}

	// Unlock write lock.
	p.pairToMinGasPriceMutex.Unlock()
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

	// Send transaction to channel.
	p.tradeCh <- TradeAction{
		BlockNumber:         prevBlock,
		Transaction:         signedTx,
		ReplacedTransaction: tx,
		TradeOption:         option,
	}

	// Broadcast message.
	err = variables.Hub.BroadcastMsg(
		fmt.Sprintf("Transaction's gas price got updated! (%.3f Gwei)", utils.WeiToGwei(targetGasPrice)),
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

	// Send transaction to channel.
	p.tradeCh <- TradeAction{
		BlockNumber:         prevBlock,
		Transaction:         signedTx,
		ReplacedTransaction: tx,
		TradeOption:         option,
	}

	// Broadcast message.
	err = variables.Hub.BroadcastMsg("Transaction got canceled because It's not profitable anymore!")
	if err != nil {
		logger.Log.WithError(err).Errorln("Unable to broadcast message.")
	}

	return signedTx
}

// quickSortCircles
//	Quick sorts the circles.
func (p *PairUpdater) quickSortCircles() ([]*circle.TradeOption, []*big.Int) {
	// Get circles as array.
	tradeArr := make([]interface{}, 0)
	for _, value := range p.Circles {
		// Calculate optimal in.
		optimalIn, amountsOut, reserves, err := p.GetOptimalIn(value)
		if err != nil && err != variables.NoArbitrage {
			logger.Log.WithError(err).Fatalln("Unable to sort trade options.")
		} else if err == variables.NoArbitrage {
			continue
		}

		// Generate new trade option.
		option, err := circle.NewTradeOption(value, optimalIn, amountsOut, reserves)
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
	gasPrices := make([]*big.Int, len(tradeArr))
	for i, v := range tradeArr {
		tmp[i] = v.(*circle.TradeOption)
		gasPrices[i] = p.GetGasPriceForPairs(tmp[i].Circle.PairAddresses)
	}
	return tmp, gasPrices

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

			// Copy old variables.
			copy(newPath, params.Path)
			copy(newPathSymbols, params.Symbols)
			copy(newRoute, params.Route)
			copy(newRouteFees, params.RouteFees)
			copy(newRouteTokens, params.RouteTokens)

			// Append new variables.
			newPath = append(newPath, tempOutToken)
			newPathSymbols = append(newPathSymbols, tempOutSymbol)
			newRouteFees = append(newRouteFees, routeFee)
			newRoute = append(newRoute, _pair.Address())
			newRouteTokens = append(newRouteTokens, pairTokens)

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
				Path:        make([]common.Address, len(params.Path)),
				Symbols:     make([]string, len(params.Symbols)),
				Route:       make([]common.Address, len(params.Route)),
				RouteFees:   make([]*big.Int, len(params.Route)),
				RouteTokens: make([][]common.Address, len(params.Route)),
			}
			copy(newParams.Path, params.Path)
			copy(newParams.Symbols, params.Symbols)
			copy(newParams.Route, params.Route)
			copy(newParams.RouteFees, params.RouteFees)
			copy(newParams.RouteTokens, params.RouteTokens)

			newParams.Path = append(newParams.Path, tempOutToken)
			newParams.Symbols = append(newParams.Symbols, tempOutSymbol)
			newParams.Route = append(params.Route, _pair.Address())
			newParams.RouteFees = append(params.RouteFees, routeFee)
			newParams.RouteTokens = append(params.RouteTokens, pairTokens)

			// Recursive
			wg.Add(1)
			go dfsUtilOnlyCircle(newParams, resultsCh, mutex, wg, u)
		}
	}
}

// newBatchCall
// 	Generates a new BatchElem from inputs.
func (p *PairUpdater) newBatchCall(
	contractAbis []abi.ABI,
	contractAddresses []common.Address,
	functionNames []string,
	functionArgs [][]interface{},
	blockNumber *big.Int,
) (rpc.BatchElem, error) {
	// Iterate through the addresses.
	var calls []abis.MulticallCall
	for i, contractAddr := range contractAddresses {
		// Create new empty byte array.
		inputBytes, err := contractAbis[i].Pack(functionNames[i], functionArgs[i]...)
		if err != nil {
			return rpc.BatchElem{}, err
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
		return rpc.BatchElem{}, err
	}

	// Message map.
	callBlockStr := fmt.Sprintf("0x%x", blockNumber.Uint64())
	msg := map[string]interface{}{
		"to":   p.params.Multicaller.Address.String(),
		"data": fmt.Sprintf("0x%s", hex.EncodeToString(callBytes)),
		"gas":  fmt.Sprintf("0x%x", 3_000_000),
	}

	// Response variables..
	var resHex hexutil.Bytes
	var resError error

	// Batch element.
	return rpc.BatchElem{
		Method: "eth_call",
		Args:   []interface{}{msg, callBlockStr},
		Result: &resHex,
		Error:  resError,
	}, nil
}

// multiCallBatch
//	Batch calls the multicaller.
func (p *PairUpdater) multiCallBatch(
	contractAbis []abi.ABI,
	contractAddresses []common.Address,
	functionNames []string,
	functionArgs [][]interface{},
	blockNumber *big.Int,
) (*big.Int, [][]byte, error) {
	// Check inputs.
	if len(contractAbis) != len(contractAddresses) ||
		len(contractAbis) != len(functionNames) ||
		len(contractAbis) != len(functionArgs) {
		return nil, nil, variables.InvalidInput
	}

	// Each call limits.
	callGasUsage := 21_000
	maxGasUsage := 3_000_000

	// The chunk size. (+1 to make sure)
	chunkSize := (callGasUsage * len(contractAbis) / maxGasUsage) + 1
	chunkSize += 1 // to make sure gas limit is not exceeded.

	// Split the chunk.
	var abiChunks [][]abi.ABI = chunkBy(contractAbis, chunkSize)
	var addressChunks [][]common.Address = chunkBy(contractAddresses, chunkSize)
	var nameChunks [][]string = chunkBy(functionNames, chunkSize)
	var argsChunks [][][]interface{} = chunkBy(functionArgs, chunkSize)

	// Check chunk length.
	if len(abiChunks) != len(addressChunks) ||
		len(abiChunks) != len(nameChunks) ||
		len(abiChunks) != len(argsChunks) ||
		len(abiChunks) != chunkSize {

		fmt.Println(len(abiChunks), len(addressChunks), len(nameChunks), len(argsChunks), chunkSize)
		panic("something is wrong with the chunk sizes")
	}

	var err error

	// Check block number.
	if blockNumber.Cmp(common.Big0) == 0 {
		blockNumber = new(big.Int).SetUint64(p.GetHighestBlockNumber())
	}

	// Ready multicall arguments.
	allBatchElements := make([]rpc.BatchElem, chunkSize)
	for i, _ := range addressChunks {
		allBatchElements[i], err = p.newBatchCall(abiChunks[i], addressChunks[i], nameChunks[i], argsChunks[i], blockNumber)
		if err != nil {
			return nil, nil, err
		}
	}

	// Parse batch elements into chunks.
	var chunkCount = (len(allBatchElements) / 2) + 1
	var batchElementChunks [][]rpc.BatchElem = chunkBy(allBatchElements, chunkCount)

	// Wait group and channel.
	wg := new(sync.WaitGroup)
	ch := make(chan struct {
		Elements []rpc.BatchElem
		Index    int
	}, chunkCount)

	// Batch call for each chunk.
	for i, batchElementChunk := range batchElementChunks {
		// Increase the counter.
		wg.Add(1)

		// Start batch call in another goroutine.
		go func(chunk []rpc.BatchElem, index int) {
			defer wg.Done()

			// Batch call.
			err = p.rpcBackend.BatchCallContext(context.Background(), chunk)
			if err != nil {
				logger.Log.WithError(err).Fatalln("Unable to batch call.")
			}

			// Send to the channel.
			ch <- struct {
				Elements []rpc.BatchElem
				Index    int
			}{
				Elements: chunk,
				Index:    index,
			}
		}(batchElementChunk, i)
	}

	// Wait.
	wg.Wait()
	close(ch)

	// The return bytes map.
	chunksMap := make(map[int][][]byte)

	// Iterate over channel.
	for tmp := range ch {
		// Output response.
		var batchReturnBytes = make([][]byte, 0)

		// Iterate over elements
		for _, batchElement := range tmp.Elements {
			// Check error.
			if batchElement.Error != nil {
				return nil, nil, batchElement.Error
			}

			// Check result.
			if batchElement.Result == nil {
				return nil, nil, variables.EmptyResponse
			}

			// Unpack the result.
			var out, err = multicallerAbi.Unpack("aggregate", *(batchElement.Result.(*hexutil.Bytes)))
			if err != nil {
				return nil, nil, err
			}

			// The call block number.
			callBlockNum := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			if blockNumber.Cmp(common.Big0) != 0 && callBlockNum.Cmp(blockNumber) != 0 {
				panic("call block number is not right")
			}

			// The response bytes.
			resBytes := out[1].([][]byte)
			batchReturnBytes = append(batchReturnBytes, resBytes...)
		}

		// Append to the map.
		chunksMap[tmp.Index] = batchReturnBytes
	}

	// Order the responses and combine into one.
	allBatchReturnBytes := make([][]byte, 0)
	for i := 0; i < chunkCount; i++ {
		allBatchReturnBytes = append(allBatchReturnBytes, chunksMap[i]...)
	}

	// Check output.
	if len(contractAddresses) != len(allBatchReturnBytes) {
		panic("missing multicall response")
	}

	return blockNumber, allBatchReturnBytes, nil
}

// chunkBy
// 	Splits items into chunks.
func chunkBy[T any](items []T, size int) (chunks [][]T) {
	// Check parameters.
	if len(items) < size {
		panic("chunk size too much")
	}

	// The chunks.
	chunks = make([][]T, size)

	// Chunk size.
	chunkSize := len(items) / size
	lastChunkSize := chunkSize + (len(items) % size)

	// Iterate over chunks.
	var itemCur = 0
	for chunkId := 0; chunkId < size-1; chunkId++ {
		chunks[chunkId] = items[itemCur : itemCur+chunkSize]
		itemCur = itemCur + chunkSize
	}

	// The last chunk.
	chunks[size-1] = items[itemCur : itemCur+lastChunkSize]
	return chunks
}
