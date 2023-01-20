package variables

import "errors"

// The errors.

var InvalidInput = errors.New("invalid input")
var AlreadyStarted = errors.New("already started")
var TooManyRouters = errors.New("too many routers")
var EmptyResponse = errors.New("response is empty")
var NoFactoryFound = errors.New("no factory found")
var TooManyPairs = errors.New("too many pairs")
var NoPairFound = errors.New("no pair found")
var NoArbitrage = errors.New("no arbitrage change")
var InsufficientLiquidity = errors.New("insufficient liquidity")
var OverFlow = errors.New("reserve overflow")

var UnableToUnlock = errors.New("unable to unlock account")
