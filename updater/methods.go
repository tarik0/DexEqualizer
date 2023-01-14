package updater

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/config"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/utils"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
	"time"
)

// Start
//	Finds the pair addresses from the routers and start's listening.
func (p *PairUpdater) Start() error {
	// Don't start if already.
	if len(p.Factories) > 0 {
		return variables.AlreadyStarted
	}

	// Find factories.
	var err error
	err = p.findFactories()
	if err != nil {
		return err
	}

	// Find pair addresses.
	err = p.findPairAddresses()
	if err != nil {
		return err
	}

	// Find pair reserves.
	err = p.findReserves()
	if err != nil {
		return err
	}

	// Find token decimals.
	err = p.findDecimals()
	if err != nil {
		return err
	}

	// Find pair circles.
	err = p.findCircles()
	if err != nil {
		return err
	}

	// Sort circles.
	p.sortCircles()

	// Subscribe to new blocks.
	p.subscribeToHeads()

	// Subscribe to events.
	p.subscribeToSync()

	// Get current block number.
	blockNum, err := p.backend.BlockNumber(context.Background())
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to get block number.")
	}
	p.lastBlockNum.Store(blockNum)

	// Start listening for events.
	go func() {
		var err error
		var vLog types.Log

		for {
			select {
			case err = <-p.logsSub.Err():
				// Disconnected, retry.
				close(p.logsCh)
				logger.Log.WithError(err).Errorln("Disconnected from the logs! Reconnecting...")
				p.subscribeToSync()
				logger.Log.WithError(err).Errorln("Connected back to the logs!")
			case vLog = <-p.logsCh:
				// Redirect to listen method.
				go p.listenSync(vLog)
			}
		}
	}()

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
					go p.listenBlocks(header)
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

	return pairFee, nil
}

// GetSortedTrades
//	Returns sorted options.
func (p *PairUpdater) GetSortedTrades() []*circle.TradeOption {
	val := p.sortedTrades.Load()
	if val == nil {
		return nil
	}

	return val.([]*circle.TradeOption)
}

// GetSortTime
//	Returns the sort time.
func (p *PairUpdater) GetSortTime() time.Duration {
	val := p.sortTime.Load()
	if val == nil {
		return time.Duration(0)
	}

	return val.(time.Duration)
}

// GetBlockNumber
//	Returns the latest block number.
func (p *PairUpdater) GetBlockNumber() uint64 {
	val := p.lastBlockNum.Load()
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
		inFee, ok := p.params.Tokens.Fees[c.Path[pairId]]
		outFee, ok := p.params.Tokens.Fees[c.Path[pairId+1]]
		if !ok {
			return nil, nil, variables.InvalidInput
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
			firstTokenFee, ok := p.params.Tokens.Fees[c.Path[0]]
			if !ok {
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

	// Limit roots.
	if rootOne.Cmp(variables.MaxOptimalIn) > 0 {
		rootOne.Set(big.NewInt(1e18))
	}
	if rootTwo.Cmp(variables.MaxOptimalIn) > 0 {
		rootTwo.Set(big.NewInt(1e18))
	}

	// Calculate amounts out.
	amountOutsOne, errOne := p.GetAmountsOut(rootOne, c.Path, c.PairAddresses)
	amountOutsTwo, errTwo := p.GetAmountsOut(rootTwo, c.Path, c.PairAddresses)

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

	// Check limit.
	if err != nil && bestAmountIn.Cmp(utils.EthersToWei(config.Parsed.ArbitrageOptions.Limiters.MaxAmountIn)) > 0 {
		bestAmountIn.Set(utils.EthersToWei(config.Parsed.ArbitrageOptions.Limiters.MaxAmountIn))
		bestAmountOut, _ = p.GetAmountsOut(bestAmountIn, c.Path, c.PairAddresses)
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
	amountsOut := []*big.Int{amountIn}

	// Iterate over pairs.
	for i, pairAddr := range route {
		// Find reserve in and out.
		resIn, resOut, err := p.AddressToPair[pairAddr].GetSortedReserves(path[i])
		if err != nil {
			return nil, err
		}

		// Get token fees.
		inputTokenFee, ok := p.params.Tokens.Fees[path[i]]
		outputTokenFee, ok := p.params.Tokens.Fees[path[i+1]]
		if !ok {
			return nil, err
		}

		// Get pair fee.
		pairFee, err := p.GetPairFee(pairAddr)
		if err != nil {
			return nil, err
		}

		// Amount in.
		tmpIn := new(big.Int).Set(amountsOut[len(amountsOut)-1])
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

		amountsOut = append(amountsOut, amountOut)
	}

	return amountsOut, nil
}
