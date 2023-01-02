package circle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/dexpair"
	"math/big"
)

// Circle
//	An arbitrage circle.
type Circle struct {
	// Path.
	Path    []common.Address
	Symbols []string

	// Pairs.
	Pairs         []*dexpair.DexPair
	PairFees      []*big.Int
	PairTokens    [][]common.Address
	PairAddresses []common.Address
}

// TradeOption
// 	Is a profitable arbitrage path.
type TradeOption struct {
	Circle *Circle

	OptimalIn  *big.Int
	AmountsOut []*big.Int
}

// TradeOptionJSON
//	TradeOption's json representation.
type TradeOptionJSON struct {
	TriggerLimit *big.Int `json:"TriggerLimit"`
	Pairs        []string `json:"Pairs"`
	Path         []string `json:"Path"`
	Symbols      []string `json:"Symbols"`
	AmountsOut   []string `json:"AmountsOut"`
}
