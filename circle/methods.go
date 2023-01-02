package circle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
	"strings"
)

// ID is the unique pair hash.
func (c *Circle) ID() uint64 {
	return CalculateID(c.PairAddresses)
}

// SymbolsStr is the path symbols as string.
func (c *Circle) SymbolsStr() string {
	return strings.Join(c.Symbols, "->")
}

// GetProfit returns the profit.
func (t *TradeOption) GetProfit() (*big.Int, error) {
	if t.Circle.Path[0] != t.Circle.Path[len(t.Circle.Path)-1] {
		return nil, variables.InvalidInput
	}

	// Calculate profit.
	profit := new(big.Int).Sub(
		t.AmountsOut[len(t.AmountsOut)-1],
		t.AmountsOut[0],
	)

	// Calculate loan debt.
	/*loanDebt := new(big.Int).Mul(
		new(big.Int).Sub(variables.Big10000, t.Circle.PairFees[0]),
		t.AmountsOut[0],
	)
	loanDebt.Div(loanDebt, variables.Big10000)

	// Subs profit.
	return profit.Sub(profit, loanDebt), nil
	*/

	// Subs profit.
	return profit, nil
}

// TriggerLimit returns the trigger limit.
func (t *TradeOption) TriggerLimit() *big.Int {
	// Trigger limit.
	gasPrice := new(big.Int).Mul(big.NewInt(5), big.NewInt(1e9)) // 5Gwei
	gasLimit := big.NewInt(int64(len(t.Circle.Pairs) * 150000))
	gasCost := new(big.Int).Mul(gasPrice, gasLimit)

	return gasCost
}

func (t *TradeOption) GetJSON() TradeOptionJSON {
	// Path.
	var pathStr = make([]string, len(t.Circle.Path))
	for i, val := range t.Circle.Path {
		pathStr[i] = val.String()
	}

	// Route.
	var pairsStr = make([]string, len(t.Circle.Pairs))
	for i, val := range t.Circle.Pairs {
		pairsStr[i] = val.Address.String()
	}

	// Amounts.
	var amountsStr = make([]string, len(t.AmountsOut))
	for i, val := range t.AmountsOut {
		amountsStr[i] = val.String()
	}

	return TradeOptionJSON{
		Path:         pathStr,
		Symbols:      t.Circle.Symbols,
		Pairs:        pairsStr,
		AmountsOut:   amountsStr,
		TriggerLimit: t.TriggerLimit(),
	}
}

// CalculateID calculates pair id.
func CalculateID(route []common.Address) uint64 {
	hash := fnv1a.Init64
	for _, pairAddr := range route {
		hash = fnv1a.AddBytes64(hash, pairAddr.Bytes())
	}
	return hash
}
