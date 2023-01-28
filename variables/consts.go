package variables

import (
	"math/big"
)

// The constant variables.

var Big10000 = big.NewInt(10000)
var MaxUInt112, _ = big.NewInt(0).SetString("5192296858534827628530496329220095", 10)

var GasPrice = big.NewInt(7e9)
var ChiCost = big.NewInt(32068600000000)
