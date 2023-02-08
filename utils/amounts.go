package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/tarik0/DexEqualizer/variables"
	"math/big"
)

// GetAmountOut is the same getAmountOut but with reserves.
func GetAmountOut(
	amountIn *big.Int,
	pairFee *big.Int,
	reserveIn *big.Int,
	reserveOut *big.Int,
) (*big.Int, *big.Int, error) {
	if amountIn.Cmp(common.Big0) <= 0 {
		return nil, nil, variables.InvalidInput
	}
	if reserveIn.Cmp(common.Big0) <= 0 || reserveOut.Cmp(common.Big0) <= 0 {
		return nil, nil, variables.InsufficientLiquidity
	}
	if new(big.Int).Add(amountIn, reserveIn).Cmp(variables.MaxUInt112) >= 0 {
		return nil, nil, variables.OverFlow
	}

	amountInWithFee := new(big.Int).Mul(amountIn, pairFee)
	num := new(big.Int).Mul(reserveOut, amountInWithFee)
	den := new(big.Int).Add(new(big.Int).Mul(reserveIn, variables.Big10000), amountInWithFee)
	amountOut := big.NewInt(0).Div(num, den)

	if amountOut.Cmp(common.Big0) <= 0 {
		amountOut = new(big.Int).Set(common.Big0)
	}

	return amountIn, amountOut, nil
}

// GetAmountIn is the same getAmountIn but with reserves.
func GetAmountIn(
	amountOut *big.Int,
	pairFee *big.Int,
	reserveIn *big.Int,
	reserveOut *big.Int,
) (*big.Int, *big.Int, error) {
	if amountOut.Cmp(common.Big0) <= 0 {
		return nil, nil, variables.InvalidInput
	}
	if reserveIn.Cmp(common.Big0) <= 0 || reserveOut.Cmp(common.Big0) <= 0 {
		return nil, nil, variables.InsufficientLiquidity
	}

	// uint numerator = reserveIn.mul(amountOut).mul(1000);
	// uint denominator = reserveOut.sub(amountOut).mul(997);
	// amountIn = (numerator / denominator).add(1);

	num := new(big.Int).Mul(reserveIn, amountOut)
	num.Mul(num, variables.Big10000)

	den := new(big.Int).Sub(reserveOut, amountOut)
	den.Mul(den, pairFee)

	amountIn := num.Div(num, den)
	amountIn.Add(amountIn, common.Big1)

	// Check if amount in is below zero.
	if amountIn.Cmp(common.Big0) <= 0 {
		amountIn = new(big.Int).Set(common.Big0)
	}

	return amountIn, amountOut, nil
}

// CutFee cuts fee from the token amount.
func CutFee(
	amountIn *big.Int,
	fee *big.Int,
) *big.Int {
	return new(big.Int).Div(
		new(big.Int).Mul(amountIn, fee),
		variables.Big10000,
	)
}
