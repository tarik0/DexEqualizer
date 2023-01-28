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
	return strings.Join(c.Symbols, "->")
}

// NormalProfit returns the profit.
func (t *TradeOption) NormalProfit() (*big.Int, error) {
	if !bytes.Equal(t.Circle.Path[0].Bytes(), t.Circle.Path[len(t.Circle.Path)-1].Bytes()) {
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

// LoanProfit returns the flashloan profit.
func (t *TradeOption) LoanProfit() (*big.Int, error) {
	// Get normal profit.
	normalProfit, err := t.NormalProfit()
	if err != nil {
		return nil, err
	}

	// Subtract the loan debt.
	normalProfit.Sub(normalProfit, t.LoanDebt())
	return normalProfit, nil
}

// LoanDebt returns the flashloan debt.
func (t *TradeOption) LoanDebt() *big.Int {
	// Calculate loan debt.
	loanDebt := new(big.Int).Mul(
		new(big.Int).Sub(variables.Big10000, t.Circle.PairFees[0]),
		t.AmountsOut[0],
	)
	loanDebt.Div(loanDebt, variables.Big10000)

	return loanDebt
}

// LoanTriggerProfit returns the minimum amount of profit we need to trigger a flashloan swap.
func (t *TradeOption) LoanTriggerProfit(gasPrice *big.Int) *big.Int {
	gasCost := new(big.Int).SetUint64(t.LoanGasSpent() - t.LoanChiRefund(gasPrice))
	gasCost.Mul(gasCost, gasPrice)

	// Add chi cost.
	chiCost := new(big.Int).Mul(new(big.Int).SetUint64(t.LoanGasTokenAmount()), variables.ChiCost)
	gasCost.Add(gasCost, chiCost)

	// Add loan cost.
	gasCost.Add(gasCost, t.LoanDebt())
	return gasCost
}

// LoanChiRefund is the amount of gas that's going to get refunded.
func (t *TradeOption) LoanChiRefund(gasPrice *big.Int) uint64 {
	gasRatio := gasPrice.Uint64() / 5e9
	efficiency := (24000 * gasRatio / (35678 + 6053*gasRatio)) * 100
	refund := t.LoanGasTokenAmount() * 16000 * efficiency / 100
	if refund > (t.LoanGasSpent() / 2) {
		return t.LoanGasSpent() / 2
	}
	return refund
}

// LoanGasSpent returns the gas spent for the circle.
func (t *TradeOption) LoanGasSpent() uint64 {
	// Gas spent.
	var gasSpent uint64 = 21000 // initialize gas.

	// Calculate message data length.
	msgDataLength := uint64(len(t.Circle.Pairs) * 20)     // pairs
	msgDataLength += uint64(len(t.Circle.Pairs) * 2 * 32) // reserves
	msgDataLength += uint64(len(t.Circle.Path) * 20)      // path
	msgDataLength += uint64(len(t.AmountsOut) * 32)       // amounts out
	msgDataLength += 20                                   // gas token
	msgDataLength += 32                                   // gas token amount
	msgDataLength += 32                                   // pool debt
	msgDataLength += 1                                    // revert on reserve change
	gasSpent += 16 * msgDataLength

	// Get reserves call
	gasSpent += uint64(len(t.Circle.Pairs) * 5000) // 5k gas each call.

	// Swap fees.
	gasSpent += uint64(len(t.Circle.Pairs) * 94000) // 94K gas each swap.

	// The burn tokens call.
	gasSpent += 21000

	return gasSpent
}

// LoanGasTokenAmount returns the gas token amount to get used on loan swap.
func (t *TradeOption) LoanGasTokenAmount() uint64 {
	gasTokens := uint64((t.LoanGasSpent() + 14154) / 41947)
	return gasTokens
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
		TriggerLimit: t.LoanTriggerProfit(variables.GasPrice),
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
