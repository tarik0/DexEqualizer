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

	// Subs profit.
	return profit, nil
}

// Gas returns the transaction gas.
func (t *TradeOption) Gas() uint64 {
	return uint64(len(t.Circle.Pairs)) * variables.GasPerHopChi
}

// TriggerProfit returns the minimum amount of profit we need to trigger a swap.
func (t *TradeOption) TriggerProfit() *big.Int {
	gasCost := new(big.Int).SetUint64(t.TriggerGas())
	gasCost.Mul(gasCost, variables.GasPrice)

	return gasCost
}

// TriggerGas returns the maximum amount of gas this circle should use to profit.
func (t *TradeOption) TriggerGas() uint64 {
	gas := uint64(len(t.Circle.Pairs)) * variables.GasPerHop
	gas -= uint64(len(t.Circle.Pairs)+1) * variables.ChiRefundGas
	gas -= uint64(len(t.Circle.Pairs)-2) * 5000
	return gas
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
		TriggerLimit: t.TriggerProfit(),
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
