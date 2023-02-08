package dexpair

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/logger"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
	"sync/atomic"
)

// Address is the pair address.
func (d *DexPair) Address() common.Address {
	return d.address
}

// TokenA is the A token address.
func (d *DexPair) TokenA() common.Address {
	return d.tokenA
}

// TokenB is the second token address.
func (d *DexPair) TokenB() common.Address {
	return d.tokenB
}

// SetReserves updates the pair reserves.
func (d *DexPair) SetReserves(reserveA *big.Int, reserveB *big.Int, blockNum *big.Int, txIndex *big.Int, logIndex *big.Int) {
	// Check concurrency.
	if old := atomic.SwapInt32(d.isConcurrent, 1); old == 1 {
		panic("concurrent reserve read/write")
	}
	defer atomic.StoreInt32(d.isConcurrent, 0)

	// The flags.
	compareBlock := blockNum.Cmp(d.lastUpdateBlock)
	compareTxIndex := txIndex.Cmp(d.lastUpdateTxIndex)
	compareLogIndex := logIndex.Cmp(d.lastUpdateLogIndex)

	// Return if block is lower.
	if compareBlock < 0 {
		return
	}

	// Return if block is same but tx index is lower.
	if compareBlock == 0 && compareTxIndex < 0 {
		return
	}

	// Return if block and the tx index is same but log index is lower.
	if compareBlock == 0 && compareTxIndex == 0 && compareLogIndex < 0 {
		return
	}

	// Return if it's the same log.
	if compareBlock == 0 && compareTxIndex == 0 && compareLogIndex == 0 {
		return
	}

	logger.Log.
		WithField("address", d.address).
		WithField("res0", reserveA).
		WithField("res1", reserveB).
		WithField("block", blockNum).
		WithField("txIndex", txIndex).
		WithField("logIndex", logIndex).
		Debugln("Pair reserves updated!")

	d.res0.Set(reserveA)
	d.res1.Set(reserveB)
	d.lastUpdateTxIndex.Set(txIndex)
	d.lastUpdateBlock.Set(blockNum)

	// Compare variables.
	if d.res0.Cmp(reserveA) != 0 || d.res1.Cmp(reserveB) != 0 {
		panic("wtf")
	}
}

// GetReserves returns the reserves.
func (d *DexPair) GetReserves() []*big.Int {
	// Check concurrency.
	if old := atomic.SwapInt32(d.isConcurrent, 1); old == 1 {
		panic("concurrent reserve read/write")
	}
	defer atomic.StoreInt32(d.isConcurrent, 0)

	return []*big.Int{new(big.Int).Set(d.res0), new(big.Int).Set(d.res1)}
}

// GetSortedReserves sorts the reserves and returns.
func (d *DexPair) GetSortedReserves(address common.Address) (reserveIn *big.Int, reserveOut *big.Int, err error) {
	// Check concurrency.
	if old := atomic.SwapInt32(d.isConcurrent, 1); old == 1 {
		panic("concurrent reserve read/write")
	}
	defer atomic.StoreInt32(d.isConcurrent, 0)

	// Sort reserves.
	if address == d.tokenA {
		reserveIn, reserveOut, err = new(big.Int).Set(d.res0), new(big.Int).Set(d.res1), nil
	} else if address == d.tokenB {
		reserveIn, reserveOut, err = new(big.Int).Set(d.res1), new(big.Int).Set(d.res0), nil
	} else {
		reserveIn, reserveOut, err = nil, nil, variables.InvalidInput
	}

	return reserveIn, reserveOut, err
}

// SortTokens
//	Sorts the tokens like Uniswap.
func SortTokens(tokenA common.Address, tokenB common.Address) []common.Address {
	var tmp = make([]common.Address, 2)

	a := new(big.Int).SetBytes(tokenA.Bytes())
	b := new(big.Int).SetBytes(tokenB.Bytes())

	if a.Cmp(b) < 0 {
		tmp[0], tmp[1] = common.BytesToAddress(a.Bytes()), common.BytesToAddress(b.Bytes())
	} else {
		tmp[1], tmp[0] = common.BytesToAddress(a.Bytes()), common.BytesToAddress(b.Bytes())
	}

	return tmp
}
