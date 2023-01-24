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

// NormalGas returns the transaction gas for a swap.
func (t *TradeOption) NormalGas() uint64 {
	swapCost := uint64(len(t.Circle.Pairs)) * variables.NormalGasPerHop
	return swapCost + variables.LoanChiBurnCost
}

// LoanGas returns the transaction gas for a flashloan swap.
func (t *TradeOption) LoanGas() uint64 {
	swapCost := uint64(len(t.Circle.Pairs)) * (variables.LoanGasPerHop + 35000)
	return swapCost + variables.LoanChiBurnCost
}

// NormalTriggerGas returns the maximum amount of gas this circle should use to profit.
func (t *TradeOption) NormalTriggerGas() uint64 {
	gas := uint64(len(t.Circle.Pairs)) * variables.NormalGasPerHop
	gas += variables.NormalChiBurnCost
	gas -= t.NormalGasTokenAmount() * variables.NormalChiRefundGas
	return gas
}

// LoanTriggerGas returns the maximum amount of gas this circle should use to profit.
func (t *TradeOption) LoanTriggerGas() uint64 {
	gas := uint64(len(t.Circle.Pairs)) * variables.LoanGasPerHop
	gas += variables.LoanChiBurnCost
	gas -= t.LoanGasTokenAmount() * variables.LoanChiRefundGas
	return gas
}

// NormalTriggerProfit returns the minimum amount of profit we need to trigger a swap.
func (t *TradeOption) NormalTriggerProfit() *big.Int {
	gasCost := new(big.Int).SetUint64(t.NormalTriggerGas())
	gasCost.Mul(gasCost, variables.GasPrice)

	// Add chi cost.
	chiCost := new(big.Int).Mul(new(big.Int).SetUint64(t.NormalGasTokenAmount()), variables.ChiCost)
	gasCost.Add(gasCost, chiCost)

	return gasCost
}

// LoanTriggerProfit returns the minimum amount of profit we need to trigger a flashloan swap.
func (t *TradeOption) LoanTriggerProfit() *big.Int {
	gasCost := new(big.Int).SetUint64(t.LoanTriggerGas())
	gasCost.Mul(gasCost, variables.GasPrice)

	// Add chi cost.
	chiCost := new(big.Int).Mul(new(big.Int).SetUint64(t.LoanGasTokenAmount()), variables.ChiCost)
	gasCost.Add(gasCost, chiCost)

	// Add loan cost.
	gasCost.Add(gasCost, t.LoanDebt())

	return gasCost
}

// NormalGasTokenAmount returns the gas token amount to get used on normal swap.
func (t *TradeOption) NormalGasTokenAmount() uint64 {
	return uint64(len(t.Circle.Pairs) * 2)
}

// LoanGasTokenAmount returns the gas token amount to get used on loan swap.
func (t *TradeOption) LoanGasTokenAmount() uint64 {
	return uint64(len(t.Circle.Pairs)*2) + 1
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
		TriggerLimit: t.LoanTriggerProfit(),
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
