package dexpair

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// DexPair
//	A struct for Uniswap V2 pair.
type DexPair struct {
	// The reserves.
	res0 *big.Int
	res1 *big.Int

	// The addresses.
	tokenA common.Address
	tokenB common.Address

	// The pair address.
	address common.Address

	// Update info.
	lastUpdateBlock    *big.Int
	lastUpdateTxIndex  *big.Int
	lastUpdateLogIndex *big.Int

	// The concurrency checkers.
	isConcurrent *int32
}

// NewDexPair
//	Generates a new DexPair.
func NewDexPair(address common.Address, tokenA common.Address, tokenB common.Address) *DexPair {
	sorted := SortTokens(tokenA, tokenB)
	var isConcurrent int32 = 0
	return &DexPair{
		address:            address,
		tokenA:             sorted[0],
		tokenB:             sorted[1],
		res0:               new(big.Int).Set(common.Big0),
		res1:               new(big.Int).Set(common.Big0),
		lastUpdateBlock:    new(big.Int).Set(common.Big0),
		lastUpdateTxIndex:  new(big.Int).Set(common.Big0),
		lastUpdateLogIndex: new(big.Int).Set(common.Big0),
		isConcurrent:       &isConcurrent,
	}
}
