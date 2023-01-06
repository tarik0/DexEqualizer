package updater

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/dexpair"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
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

	// TODO: remove
	var shit int32 = 0

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

						atomic.AddInt32(&shit, 1)
						logger.Log.Infoln(fmt.Sprintf("[%d/%d] %s | %s, %s", atomic.LoadInt32(&shit), len(p.Pairs), pairAddrs[k].String(), pairReserves[k][0].String(), pairReserves[k][1].String()))
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
		MaxResultCount: MaxCircleResults,
		MaxHops:        MaxHops,
		Path:           path,
		Symbols:        pathSymbols,
		Route:          make([]common.Address, 0),
		RouteFees:      make([]*big.Int, 0),
		RouteTokens:    make([][]common.Address, 0),
	}

	// Wait group.
	var wg sync.WaitGroup

	// The results map.
	var resCount int32 = 0
	var resMutex = sync.RWMutex{}

	p.PairToCircles = make(map[common.Address]map[uint64]*circle.Circle)
	p.Circles = make(map[uint64]*circle.Circle, 0)

	// Start DFS.
	wg.Add(1)
	dfsUtilOnlyCircle(dfsParams, &p.PairToCircles, &p.Circles, &resMutex, &resCount, &wg, p)
	wg.Wait()

	return nil
}

// sortCircles
//	Sorts DFS circles.
func (p *PairUpdater) sortCircles() {
	p.sortMutex.Lock()
	p.sortedTrades = p.quickSortCircles()
	p.sortMutex.Unlock()
}

// subscribeToSync
//	Subscribes to the sync event.
func (p *PairUpdater) subscribeToSync() {
	// To subscribe query.
	query := ethereum.FilterQuery{
		Addresses: maps.Keys(p.AddressToPair),
	}

	// Make new channel.
	p.logsCh = make(chan types.Log)

	// Subscribe to the events.
	var err error
	p.logsSub, err = p.backend.SubscribeFilterLogs(context.Background(), query, p.logsCh)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to subscribe to 'sync' logs!")
	}
}

// listenSync
//	Listens new events.
func (p *PairUpdater) listenSync(vLog types.Log) {
	// Iterate over logs.
	startTime := time.Now()
	for _, log := range vLog.Topics {
		// Continue if topic is not sync.
		if !bytes.Equal(syncId.Bytes(), log.Bytes()) {
			continue
		}

		// Decode event.
		syncDetails, err := pairAbi.Unpack("Sync", vLog.Data)
		if err != nil {
			logger.Log.WithError(err).Errorln("Unable to decode 'sync' events.")
			continue
		}

		// Update reserves.
		resA, resB := syncDetails[0].(*big.Int), syncDetails[1].(*big.Int)
		p.AddressToPair[vLog.Address].SetReserves(resA, resB)
	}
	updateElapsed := time.Since(startTime)

	// Sort again.
	startTime = time.Now()
	p.sortCircles()
	sortElapsed := time.Since(startTime)

	// Run callback.
	if p.OnSync != nil {
		go p.OnSync(updateElapsed, sortElapsed, p)
	}
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

		// Append to the list.
		tradeArr = append(tradeArr, &circle.TradeOption{
			Circle:     value,
			OptimalIn:  optimalIn,
			AmountsOut: amountsOut,
		})
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
		profitOne, err := arr[j].GetProfit()
		profitTwo, err := pivot.GetProfit()
		if err != nil {
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
func dfsUtilOnlyCircle(params DFSCircleParams, table *map[common.Address]map[uint64]*circle.Circle, circles *map[uint64]*circle.Circle, resMutex *sync.RWMutex, resCount *int32, wg *sync.WaitGroup, u *PairUpdater) {
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
		if len(params.Path) > params.MaxHops {
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
		if bytes.Equal(tempOutToken.Bytes(), params.Path[0].Bytes()) {
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
			newRoute = append(newRoute, _pair.Address)
			newRouteTokens = append(newRouteTokens, pairTokens)

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
			}

			// Append to channel.
			resMutex.Lock()

			// Check limiter.
			if atomic.LoadInt32(resCount) > int32(params.MaxResultCount) {
				resMutex.Unlock()
				break
			}

			// Iterate over routes.
			for _, pairAddr := range arbCircle.PairAddresses {
				if _, ok = (*table)[pairAddr]; !ok {
					(*table)[pairAddr] = make(map[uint64]*circle.Circle, 0)
				}

				circleId := arbCircle.ID()
				(*table)[pairAddr][circleId] = arbCircle
				(*circles)[circleId] = arbCircle
			}

			// Increase counter.
			atomic.AddInt32(resCount, 1)
			resMutex.Unlock()
		} else {
			// The params.
			newParams := DFSCircleParams{
				MaxHops:        params.MaxHops,
				MaxResultCount: params.MaxResultCount,
				Path:           make([]common.Address, len(params.Path)),
				Symbols:        make([]string, len(params.Symbols)),
				Route:          make([]common.Address, len(params.Route)),
				RouteFees:      make([]*big.Int, len(params.Route)),
				RouteTokens:    make([][]common.Address, len(params.Route)),
			}
			copy(newParams.Path, params.Path)
			copy(newParams.Symbols, params.Symbols)
			copy(newParams.Route, params.Route)
			copy(newParams.RouteFees, params.RouteFees)
			copy(newParams.RouteTokens, params.RouteTokens)

			newParams.Path = append(newParams.Path, tempOutToken)
			newParams.Symbols = append(newParams.Symbols, tempOutSymbol)
			newParams.Route = append(params.Route, _pair.Address)
			newParams.RouteFees = append(params.RouteFees, routeFee)
			newParams.RouteTokens = append(params.RouteTokens, pairTokens)

			// Recursive
			wg.Add(1)
			go dfsUtilOnlyCircle(newParams, table, circles, resMutex, resCount, wg, u)
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
	gasLimit, err := p.backend.EstimateGas(context.TODO(), msg)
	result, err := p.backend.PendingCallContract(context.TODO(), msg)
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
