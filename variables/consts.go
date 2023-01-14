package variables

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

// The constant variables.

var Big10000 = big.NewInt(10000)
var MaxUInt112, _ = big.NewInt(0).SetString("5192296858534827628530496329220095", 10)
var MaxOptimalIn = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(19), nil)
var EmptyCallOpts = bind.CallOpts{}

var GasPrice = big.NewInt(13e9)
var GasPerHop = big.NewInt(87000)
