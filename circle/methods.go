package circle

import (
	"bytes"
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
	return strings.Join(c.Symbols, " -> ")
}

// PairAddressesStr is the pair addresses as string.
func (c *Circle) PairAddressesStr() string {
	var tmp = make([]string, len(c.PairAddresses))
	for i, addr := range c.PairAddresses {
		tmp[i] = addr.String()
	}
	return strings.Join(tmp, " -> ")
}

// NormalProfit returns the profit.
func (t *TradeOption) NormalProfit() (*big.Int, error) {
	if !bytes.EqualFold(t.Circle.Path[0].Bytes(), t.Circle.Path[len(t.Circle.Path)-1].Bytes()) {
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

// GetTradeCost returns trade cost of this option.
func (t *TradeOption) GetTradeCost(gasPrice *big.Int) *big.Int {
	gasCost := new(big.Int).SetUint64(t.NormalGasSpentWithBurn() - t.NormalChiRefund())
	gasCost.Mul(gasCost, gasPrice)

	// Add chi cost.
	chiCost := new(big.Int).Mul(new(big.Int).SetUint64(t.NormalGasTokenAmount()), variables.ChiCost)
	gasCost.Add(gasCost, chiCost)
	return gasCost
}

// MaxGasPrice returns the maximum gas price for this option to profit.
func (t *TradeOption) MaxGasPrice() (*big.Int, error) {
	// Get profit.
	profit, err := t.NormalProfit()
	if err != nil {
		return nil, err
	}

	// Subtract the chi cost.
	chiCost := new(big.Int).Mul(new(big.Int).SetUint64(t.NormalGasTokenAmount()), variables.ChiCost)
	profit.Sub(profit, chiCost)

	// Gas usage.
	gasCost := new(big.Int).SetUint64(t.NormalGasSpentWithBurn() - t.NormalChiRefund())

	// Calculate max gas price.
	profit.Div(profit, gasCost)
	return profit, nil
}

// NormalChiRefund is the amount of gas that's going to get refunded.
func (t *TradeOption) NormalChiRefund() uint64 {
	// ~%51 refund
	return t.NormalGasSpentWithBurn() * 51 / 100
}

// NormalGasSpent returns the gas spent for the circle.
func (t *TradeOption) NormalGasSpent() uint64 {
	// Gas spent.
	var gasSpent uint64 = 21000 // initialize gas.

	// Calculate message data length.
	msgDataLength := uint64(len(t.Circle.Pairs) * 20)     // pairs
	msgDataLength += uint64(len(t.Circle.Pairs) * 2 * 32) // reserves
	msgDataLength += uint64(len(t.Circle.Path) * 20)      // path
	msgDataLength += uint64(len(t.AmountsOut) * 32)       // amounts out
	msgDataLength += 20                                   // gas token
	msgDataLength += 32                                   // gas token amount
	msgDataLength += 1                                    // revert on reserve change
	gasSpent += 16 * msgDataLength

	// Get reserves call.
	gasSpent += uint64(len(t.Circle.Pairs) * 10000) // 5k gas each call.

	// Transfer cost.
	gasSpent += 21000

	// Swap gas cost.
	for _, tokenAddr := range t.Circle.Path[:len(t.Circle.Path)-1] {
		val, ok := variables.TargetTokens[tokenAddr]
		if !ok {
			panic("token info not found")
		}

		// Add buy transfer gas if it's not the first token.
		gasSpent += val.SwapGas.Uint64()
	}

	// The burn tokens call.
	gasSpent += 21000

	return gasSpent
}

// NormalGasTokenAmount returns the gas token amount to get used on swap.
func (t *TradeOption) NormalGasTokenAmount() uint64 {
	gasTokens := uint64((t.NormalGasSpent() + 14154) / 41947)
	return gasTokens
}

// NormalGasSpentWithBurn returns the gas spent with the gas tokens burnt.
func (t *TradeOption) NormalGasSpentWithBurn() uint64 {
	return t.NormalGasSpent() + t.NormalGasTokenAmount()*10000
}

func (t *TradeOption) GetJSON(gasPrice *big.Int) TradeOptionJSON {
	// Path.
	var pathStr = make([]string, len(t.Circle.Path))
	for i, val := range t.Circle.Path {
		pathStr[i] = val.String()
	}

	// Route.
	var pairsStr = make([]string, len(t.Circle.Pairs))
	for i, val := range t.Circle.Pairs {
		pairsStr[i] = val.Address().String()
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
		TriggerLimit: t.GetTradeCost(gasPrice),
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
