package variables

import (
	"math/big"
)

// The constant variables.

var Big10000 = big.NewInt(10000)
var MaxUInt112, _ = big.NewInt(0).SetString("5192296858534827628530496329220095", 10)

var GasPrice = big.NewInt(14e9)
var ChiCost = big.NewInt(2e13)

// Normal swap executor.

var NormalGasPerHop uint64 = 88000
var NormalChiRefundGas uint64 = 30000
var NormalChiBurnCost uint64 = 26000

// Flashloan swap executor.

var LoanGasPerHop uint64 = 127000
var LoanChiRefundGas uint64 = 32000
var LoanChiBurnCost uint64 = 30000
