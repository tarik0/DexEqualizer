package variables

import "errors"

// The errors.

var InvalidInput = errors.New("invalid input")
var AlreadyStarted = errors.New("already started")
var TooManyRouters = errors.New("too many routers")
var EmptyResponse = errors.New("response is empty")
var NoFactoryFound = errors.New("no factory found")
var InvalidBlock = errors.New("invalid block received")
var NoPairFound = errors.New("no pair found")
var NoArbitrage = errors.New("no arbitrage change")
var InsufficientLiquidity = errors.New("insufficient liquidity")
var OverFlow = errors.New("reserve overflow")
var APINotReady = errors.New("hot tokens api not ready yet")
