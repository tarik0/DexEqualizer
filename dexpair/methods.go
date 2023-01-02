package dexpair

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
)

// SetReserves updates the pair reserves.
func (d *DexPair) SetReserves(reserveA *big.Int, reserveB *big.Int) {
	d.reserveMutex.Lock()
	d.ReserveA.Set(reserveA)
	d.ReserveB.Set(reserveB)
	d.reserveMutex.Unlock()
}

// GetSortedReserves sorts the reserves and returns.
func (d *DexPair) GetSortedReserves(address common.Address) (reserveIn *big.Int, reserveOut *big.Int, err error) {
	// Sort reserves.
	if address == d.TokenA {
		reserveIn, reserveOut = d.ReserveA, d.ReserveB
	} else if address == d.TokenB {
		reserveIn, reserveOut = d.ReserveB, d.ReserveA
	} else {
		return nil, nil, variables.InvalidInput
	}
	return reserveIn, reserveOut, nil
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
