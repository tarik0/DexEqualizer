package circle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/dexpair"
	"github.com/tarik0/DexEqualizer/variables"
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

	// Is valid ?
	isValid bool
}

// TradeOption
// 	Is a profitable arbitrage path.
type TradeOption struct {
	Circle *Circle

	OptimalIn  *big.Int
	AmountsOut []*big.Int
	Reserves   [][]*big.Int

	// Is valid ?
	isValid bool
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

// NewCircle generates a new arbitrage circle.
func NewCircle(
	// Path.
	path []common.Address,
	symbols []string,

	// Pairs.
	pairs []*dexpair.DexPair,
	pairFees []*big.Int,
	pairTokens [][]common.Address,
	pairAddresses []common.Address,
) (*Circle, error) {
	// Validate the inputs.
	isValid := len(path) == len(symbols)
	isValid = len(pairs) == len(path)+1
	isValid = len(pairs) == len(pairFees)
	isValid = len(pairs) == len(pairTokens)
	isValid = len(pairs) == len(pairAddresses)
	if !isValid {
		return nil, variables.InvalidInput
	}

	return &Circle{
		path,
		symbols,
		pairs,
		pairFees,
		pairTokens,
		pairAddresses,
		true,
	}, nil
}

// NewTradeOption generates a new trade option.
func NewTradeOption(
	c *Circle,
	optimalIn *big.Int,
	amountOut []*big.Int,
	reserves [][]*big.Int,
) (*TradeOption, error) {
	// Validate inputs.
	isValid := c.isValid && optimalIn.Cmp(common.Big0) > 0
	isValid = len(amountOut) == len(c.Path)
	if !isValid {
		return nil, variables.InvalidInput
	}

	return &TradeOption{c, optimalIn, amountOut, reserves, true}, nil
}
