package updater

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/config"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"golang.org/x/exp/maps"
	"math/big"
	"sync"
)

// Start
//	Finds the pair addresses from the routers and start's listening.
func (p *PairUpdater) Start() error {
	// Don't start if already.
	if len(p.Factories) > 0 {
		return variables.AlreadyStarted
	}

	// Get current block number.
	blockNum, err := p.backend.BlockNumber(context.Background())
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get block number.")
	}
	p.highestBlockNum.Store(blockNum)
	p.lastSyncBlockNum.Store(blockNum)

	// Find factories.
	err = p.findFactories()
	if err != nil {
		return err
	}

	// Find token decimals.
	err = p.findDecimals()
	if err != nil {
		return err
	}

	// Find pair addresses.
	err = p.findPairAddresses()
	if err != nil {
		return err
	}

	// Map keys.
	p.PairAddresses = maps.Keys(p.AddressToPair)

	// Find pair reserves.
	err = p.findReserves()
	if err != nil {
		return err
	}

	// Find pair circles.
	err = p.findCircles()
	if err != nil {
		return err
	}

	// History listener.
	p.listenHistory()

	// Sort circles.
	p.sortCircles()

	// Subscribe to new blocks.
	p.subscribeToHeads()

	// Subscribe to new pending transactions.
	p.subscribeToPending()

	// The filter logs mutex.
	p.filterLogsMutex = sync.RWMutex{}

	// Start listening for new heads.
	go func() {
		var err error
		for {
			select {
			case err = <-p.blocksSub.Err():
				// Disconnected, retry.
				close(p.blocksCh)
				logger.Log.WithError(err).Errorln("Disconnected from the new blocks! Reconnecting...")
				p.subscribeToHeads()
				logger.Log.WithError(err).Errorln("Connected back to the new blocks!")
			case header := <-p.blocksCh:
				// Redirect to the listen method.
				if header != nil {
					// Update block number.
					p.TxHistoryReset <- true
					p.highestBlockNum.Store(header.Number.Uint64())
					go p.listenBlocks(header)
				}
			}
		}
	}()

	// Sometimes the pending transactions can have same account and nonce
	// so the most profitable one for the miner will get selected.
	p.accountToPendingTx = sync.Map{}

	// Start listening for new transactions.
	go func() {
		// Skip if node not supported.
		if p.pendingSub == nil || p.pendingCh == nil {
			return
		}

		for {
			select {
			case err = <-p.pendingSub.Err():
				// Disconnected, retry.
				close(p.pendingCh)
				logger.Log.WithError(err).Errorln("Disconnected from the new pending transactions! Reconnecting...")
				p.subscribeToPending()
				logger.Log.WithError(err).Errorln("Connected back to the new pending transactions!")
			case hash := <-p.pendingCh:
				if hash != nil {
					go p.listenPending(hash)
				}
			}
		}
	}()

	return nil
}

// GetPairFee
// 	Helper function to retrieve pair's fee.
func (p *PairUpdater) GetPairFee(addr common.Address) (*big.Int, error) {
	// Get pair's factory.
	pairFactory, ok := p.PairToFactory[addr]
	if !ok {
		return nil, variables.InvalidInput
	}

	// Get pair's router.
	pairRouter, ok := p.FactoryToRouter[pairFactory]
	if !ok {
		return nil, variables.InvalidInput
	}

	// Get pair fee.
	pairFee, ok := p.params.Routers.Fees[pairRouter]
	if !ok {
		return nil, variables.InvalidInput
	}

	return new(big.Int).Set(pairFee), nil
}

// GetTokenFee
//	Helper function to retrieve token's fee.
func (p *PairUpdater) GetTokenFee(addr common.Address) (*big.Int, error) {
	// Get token's fee.
	inFee, ok := p.params.Tokens.Fees[addr]
	if !ok {
		return nil, variables.InvalidInput
	}

	return new(big.Int).Set(inFee), nil
}

// GetLastSyncBlockNumber
//	Returns the latest block number.
func (p *PairUpdater) GetLastSyncBlockNumber() uint64 {
	val := p.lastSyncBlockNum.Load()
	if val == nil {
		return uint64(0)
	}

	return val.(uint64)
}

// GetHighestBlockNumber
//	Returns the highest block number.
func (p *PairUpdater) GetHighestBlockNumber() uint64 {
	val := p.highestBlockNum.Load()
	if val == nil {
		return uint64(0)
	}

	return val.(uint64)
}

// GetOptimalIn calculates the optimal input amount for maximum profit.
func (p *PairUpdater) GetOptimalIn(c *circle.Circle) (bestAmountIn *big.Int, bestAmountOut []*big.Int, err error) {
	// Check if it's a circle.
	if c.Path[0] != c.Path[len(c.Path)-1] {
		return nil, nil, variables.InvalidInput
	}

	// Check route count.
	if len(c.Pairs) < 2 {
		return nil, nil, variables.InvalidInput
	}

	// The calculation variables.
	a := new(big.Int).Set(common.Big0)
	b := new(big.Int).Set(common.Big0)
	_c := new(big.Int).Set(common.Big0)
	d := new(big.Int).Set(common.Big0)

	// Iterate over pairs.
	for pairId, pairAddr := range c.PairAddresses {
		// Get pair fee.
		pairFee, err := p.GetPairFee(pairAddr)
		if err != nil {
			return nil, nil, err
		}

		// Get token fees.
		inFee, err := p.GetTokenFee(c.Path[pairId])
		outFee, err := p.GetTokenFee(c.Path[pairId+1])
		if err != nil {
			return nil, nil, err
		}

		// Sort reserves.
		resIn, resOut, err := p.AddressToPair[pairAddr].GetSortedReserves(c.Path[pairId])
		if err != nil {
			return nil, nil, err
		}

		// Calculate A.
		if pairId == 0 {
			// _a = pairFee * inFee * outFee * resIn * resOut * _k
			a = new(big.Int).Mul(pairFee, inFee)
			a.Mul(a, outFee)
			a.Mul(a, resIn)
			a.Mul(a, resOut)
			a.Mul(a, variables.Big10000)
		} else {
			// _a = _a * pairFee * inFee * outFee * resIn * resOut * (_k ** 3)
			a.Mul(a, pairFee)
			a.Mul(a, inFee)
			a.Mul(a, outFee)
			a.Mul(a, resIn)
			a.Mul(a, resOut)
			a.Mul(a, new(big.Int).Exp(variables.Big10000, common.Big3, nil))
		}

		// Calculate B.
		if pairId == 0 {
			// _b = resIn * (_k * _k)
			b = new(big.Int).Mul(resIn, new(big.Int).Exp(variables.Big10000, common.Big2, nil))
		} else {
			// _b = _b * resIn * (_k * _k * _k)
			b.Mul(b, resIn)
			b.Mul(b, new(big.Int).Exp(variables.Big10000, common.Big3, nil))
		}

		// Calculate C.
		if pairId == 0 {
			// _c = inFee * pairFee
			_c = new(big.Int).Mul(inFee, pairFee)
		} else {
			// _c = (_c + _d) * (_k ** 3) * resIn
			_c.Mul(new(big.Int).Add(_c, d), new(big.Int).Exp(variables.Big10000, common.Big3, nil))
			_c.Mul(_c, resIn)
		}

		// Get previous output reserve and previous input reserve.
		prevResOut := new(big.Int).Set(common.Big1)
		if pairId != 0 {
			// Sort previous reserves.
			_, prevResOut, err = p.AddressToPair[c.PairAddresses[pairId-1]].GetSortedReserves(c.Path[pairId-1])
			if err != nil {
				return nil, nil, err
			}
		}

		// Calculate D.
		if pairId == 0 {
			// _d = 0
			d = new(big.Int).Set(common.Big0)
		} else if pairId == 1 {
			// Get the first token's fee.
			firstTokenFee, err := p.GetTokenFee(c.Path[0])
			if err != nil {
				return nil, nil, variables.InvalidInput
			}

			// Get the first pool's fee.
			firstPoolFee, err := p.GetPairFee(c.PairAddresses[0])
			if err != nil {
				return nil, nil, err
			}

			// _d = pairTokenFees[0][0] * pairFees[0] * inFee ** 2 * pairFee * previousOutReserve
			d = new(big.Int).Mul(firstTokenFee, firstPoolFee)
			d.Mul(d, new(big.Int).Exp(inFee, common.Big2, nil))
			d.Mul(d, pairFee)
			d.Mul(d, prevResOut)
		} else {
			// _d = _d * inFee ** 2 * pairFee * previousOutReserve
			d.Mul(d, new(big.Int).Exp(inFee, common.Big2, nil))
			d.Mul(d, pairFee)
			d.Mul(d, prevResOut)
		}
	}

	// Sqrt(a)
	a.Sqrt(a)

	// The root 2.
	num := new(big.Int).Add(a, b)
	den := new(big.Int).Add(_c, d)
	rootOne := new(big.Int).Div(num, den)
	rootOne.Abs(rootOne)

	// Root two.
	num = new(big.Int).Sub(b, a)
	rootTwo := new(big.Int).Div(num, den)
	rootTwo.Abs(rootTwo)

	// No arbitrage if one of them are zero.
	if rootOne.Cmp(common.Big0) <= 0 && rootTwo.Cmp(common.Big0) <= 0 {
		return nil, nil, variables.NoArbitrage
	}

	// Max input.
	maxIn := utils.EthersToWei(config.Parsed.ArbitrageOptions.Limiters.MaxAmountIn)

	// Limit roots.
	if rootOne.Cmp(maxIn) > 0 {
		rootOne.Set(maxIn)
	}
	if rootTwo.Cmp(maxIn) > 0 {
		rootTwo.Set(maxIn)
	}

	// Calculate amounts out.
	var errOne = variables.InvalidInput
	var errTwo = variables.InvalidInput
	var amountOutsOne []*big.Int
	var amountOutsTwo []*big.Int
	if rootTwo.Cmp(rootOne) == 0 {
		amountOutsOne, errOne = p.GetAmountsOut(rootOne, c.Path, c.PairAddresses)
	} else {
		amountOutsOne, errOne = p.GetAmountsOut(rootOne, c.Path, c.PairAddresses)
		amountOutsTwo, errTwo = p.GetAmountsOut(rootTwo, c.Path, c.PairAddresses)
	}

	// The scenarios.
	if errOne != nil && errTwo != nil {
		// No arbitrage.
		bestAmountIn, bestAmountOut, err = nil, nil, variables.NoArbitrage
	} else if errOne == nil && errTwo != nil {
		// Root one.
		bestAmountIn, bestAmountOut, err = rootOne, amountOutsOne, nil
	} else if errOne != nil && errTwo == nil {
		// Root two.
		bestAmountIn, bestAmountOut, err = rootTwo, amountOutsTwo, nil
	} else {
		// Calculate profit.
		profitOne := new(big.Int).Sub(amountOutsOne[len(amountOutsOne)-1], amountOutsOne[0])
		profitTwo := new(big.Int).Sub(amountOutsTwo[len(amountOutsTwo)-1], amountOutsTwo[0])

		if profitOne.Cmp(profitTwo) > 0 {
			bestAmountIn, bestAmountOut, err = profitOne, amountOutsOne, nil
		} else {
			bestAmountIn, bestAmountOut, err = profitTwo, amountOutsTwo, nil
		}
	}

	return bestAmountIn, bestAmountOut, err
}

// GetAmountsOut calculates amounts out.
func (p *PairUpdater) GetAmountsOut(
	amountIn *big.Int,
	path []common.Address,
	route []common.Address,
) ([]*big.Int, error) {
	// The temporary amounts out variable.
	var amountsOut = make([]*big.Int, len(path))
	amountsOut[0] = new(big.Int).Set(amountIn)

	// Iterate over pairs.
	for i, pairAddr := range route {
		// Find reserve in and out.
		resIn, resOut, err := p.AddressToPair[pairAddr].GetSortedReserves(path[i])
		if err != nil {
			return nil, err
		}

		// Get token fees.
		inputTokenFee, err := p.GetTokenFee(path[i])
		outputTokenFee, err := p.GetTokenFee(path[i+1])
		if err != nil {
			return nil, err
		}

		// Get pair fee.
		pairFee, err := p.GetPairFee(pairAddr)
		if err != nil {
			return nil, err
		}

		// Amount in.
		tmpIn := new(big.Int).Set(amountsOut[i])
		tmpIn = utils.CutFee(tmpIn, inputTokenFee)

		// Calculate
		_, amountOut, err := utils.GetAmountOut(tmpIn, pairFee, resIn, resOut)
		if err != nil {
			return nil, err
		}

		// Cut fee.
		amountOut = utils.CutFee(amountOut, outputTokenFee)

		// Check if amount out is zero.
		if amountOut.Cmp(common.Big0) <= 0 {
			return nil, variables.InvalidInput
		}

		// Append to list.
		amountsOut[i+1] = new(big.Int).Set(amountOut)
	}

	return amountsOut, nil
}
