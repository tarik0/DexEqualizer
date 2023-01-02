package dexpair

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

// DexPair
//	A struct for Uniswap V2 pair.
type DexPair struct {
	// The reserves.
	ReserveA *big.Int
	ReserveB *big.Int

	// The addresses.
	TokenA common.Address
	TokenB common.Address

	// The dexpair address.
	Address common.Address

	// Update reserve mutexes.
	reserveMutex *sync.RWMutex
}

// NewDexPair
//	Generates a new DexPair.
func NewDexPair(address common.Address, tokenA common.Address, tokenB common.Address) *DexPair {
	return &DexPair{
		Address:      address,
		TokenA:       SortTokens(tokenA, tokenB)[0],
		TokenB:       SortTokens(tokenA, tokenB)[1],
		ReserveA:     new(big.Int),
		ReserveB:     new(big.Int),
		reserveMutex: new(sync.RWMutex),
	}
}
