package variables

import (
	"math/big"
)

// The constant variables.

var Big10000 = big.NewInt(10000)
var MaxUInt112, _ = big.NewInt(0).SetString("5192296858534827628530496329220095", 10)

var CancelThresholdGasPrice = big.NewInt(51e8) // 5 Gwei
var GasPrice = big.NewInt(51e8)                // 5 GWei
var ChiCost = big.NewInt(0)
