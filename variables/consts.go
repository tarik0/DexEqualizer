package variables

import (
	"math/big"
)

// The constant variables.

var Big10000 = big.NewInt(10000)
var MaxUInt112, _ = big.NewInt(0).SetString("5192296858534827628530496329220095", 10)

var CancelThresholdGasPrice = big.NewInt(5e9 + 1) // >5 Gwei
var GasPrice = big.NewInt(5e9 + 1)                // >5 GWei
var ChiCost = big.NewInt(37341300000000.0)        // 0.0000373413 BNB
