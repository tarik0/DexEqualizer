package variables

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

// The constant variables.

var Big10000 = big.NewInt(10000)
var MaxUInt112, _ = big.NewInt(0).SetString("5192296858534827628530496329220095", 10)
var EmptyCallOpts = bind.CallOpts{}

var GasPrice = big.NewInt(13e9)
var GasPerHop uint64 = 88000
var GasPerHopChi uint64 = 130000

// ChiRefundGas is the gas that gets refunded for one Chi.
var ChiRefundGas uint64 = 10000
