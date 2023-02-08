package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/logger"
	"math/big"
)

// PrintTradeOption prints swap parameters.
func PrintTradeOption(option *circle.TradeOption) {
	logger.Log.Infoln("Pairs:")
	for i, pair := range option.Circle.PairAddresses {
		logger.Log.Infoln(fmt.Sprintf("[%d/%d] %s", i, len(option.Circle.Pairs), pair.String()))
	}
	logger.Log.Infoln("")

	logger.Log.Infoln("Path")
	for i, addr := range option.Circle.Path {
		logger.Log.Infoln(fmt.Sprintf("[%d/%d] %s", i, len(option.Circle.Path), addr.String()))
	}

	logger.Log.Infoln("")

	logger.Log.Infoln("Amounts:")
	for i, amount := range option.AmountsOut {
		logger.Log.Infoln(fmt.Sprintf("[%d/%d] %s", i, len(option.AmountsOut), amount.String()))
	}

	logger.Log.Infoln("")

	logger.Log.Infoln("Pair Tokens:")
	for i, tokens := range option.Circle.PairTokens {
		logger.Log.Infoln(fmt.Sprintf("[%d/%d] %s, %s", i, len(option.Circle.PairTokens), tokens[0].String(), tokens[1].String()))
	}

	logger.Log.Infoln("")

	logger.Log.Infoln("Reserves")
	for i, reserves := range option.Reserves {
		logger.Log.Infoln(fmt.Sprintf("[%d/%d] R0: %s, R1: %s", i, len(option.Reserves), reserves[0].String(), reserves[1].String()))
	}

	// Validate amounts.
	for i, reserves := range option.Reserves {
		var resIn = new(big.Int)
		var resOut = new(big.Int)

		if option.Circle.PairTokens[i][0].String() == option.Circle.Path[i].String() {
			resIn.Set(reserves[0])
			resOut.Set(reserves[1])
		} else {
			resIn.Set(reserves[1])
			resOut.Set(reserves[0])
		}

		// TODO take token fee.

		// Calculate amount out.
		_, calcAmountOut, err := GetAmountOut(option.AmountsOut[i], option.Circle.PairFees[i], resIn, resOut)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"pair":     option.Circle.PairAddresses[i].String(),
				"resA":     reserves[0].String(),
				"resB":     reserves[1].String(),
				"amountIn": option.AmountsOut[i].String(),
			}).WithError(err).Fatalln("Unable to calculate amounts out!")
		}

		if calcAmountOut.Cmp(option.AmountsOut[i+1]) != 0 {
			logger.Log.WithFields(logrus.Fields{
				"pair":          option.Circle.PairAddresses[i].String(),
				"resIn":         resIn.String(),
				"resOut":        resOut.String(),
				"amountIn":      option.AmountsOut[i].String(),
				"amountOut":     option.AmountsOut[i+1].String(),
				"realAmountOut": calcAmountOut.String(),
				"pairFee":       option.Circle.PairFees[i].String(),
			}).Fatalln("Amount out calculation is not right.")
		}
	}

	logger.Log.Infoln("")
	logger.Log.Infoln("Calculations are right! No problem with the parameters.")
}
