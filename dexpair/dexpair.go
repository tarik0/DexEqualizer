package dexpair

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync/atomic"
)

// DexPair
//	A struct for Uniswap V2 pair.
type DexPair struct {
	// The reserves and the latest block number.
	reservesAndLatestBlock atomic.Value // [3]*big.Int

	// The addresses.
	tokenA common.Address
	tokenB common.Address

	// The pair address.
	address common.Address
}

// NewDexPair
//	Generates a new DexPair.
func NewDexPair(address common.Address, tokenA common.Address, tokenB common.Address) *DexPair {
	p := &DexPair{
		address:                address,
		tokenA:                 SortTokens(tokenA, tokenB)[0],
		tokenB:                 SortTokens(tokenA, tokenB)[1],
		reservesAndLatestBlock: atomic.Value{},
	}
	p.SetReserves(new(big.Int).Set(common.Big0), new(big.Int).Set(common.Big0), new(big.Int).Set(common.Big0))
	return p
}
