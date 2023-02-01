package dexpair

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
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

// GetLatestUpdateBlock returns the latest block that reserves are updated.
func (d *DexPair) GetLatestUpdateBlock() (*big.Int, error) {
	val := d.reservesAndLatestBlock.Load()
	if val == nil || (val.([3]*big.Int))[2] == nil {
		return nil, variables.InvalidInput
	}
	return (val.([3]*big.Int))[2], nil
}

// SetReserves updates the pair reserves.
func (d *DexPair) SetReserves(reserveA *big.Int, reserveB *big.Int, blockNum *big.Int) {
	// Skip if blockNum < latestUpdateBlock
	latestUpdateBlock, err := d.GetLatestUpdateBlock()
	if blockNum.Cmp(common.Big0) != 0 && err != nil && blockNum.Cmp(latestUpdateBlock) < 0 {
		return
	}

	d.reservesAndLatestBlock.Store([3]*big.Int{reserveA, reserveB, blockNum})
}

// GetReserves returns the reserves.
func (d *DexPair) GetReserves() []*big.Int {
	val := d.reservesAndLatestBlock.Load().([3]*big.Int)
	return []*big.Int{val[0], val[1]}
}

// GetSortedReserves sorts the reserves and returns.
func (d *DexPair) GetSortedReserves(address common.Address) (reserveIn *big.Int, reserveOut *big.Int, err error) {
	// Get reserves.
	val := d.reservesAndLatestBlock.Load().([3]*big.Int)

	// Sort reserves.
	if address == d.tokenA {
		reserveIn, reserveOut, err = new(big.Int).Set(val[0]), new(big.Int).Set(val[1]), nil
	} else if address == d.tokenB {
		reserveIn, reserveOut, err = new(big.Int).Set(val[1]), new(big.Int).Set(val[0]), nil
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
