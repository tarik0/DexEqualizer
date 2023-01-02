package variables

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

// The token variables.

var TargetRouters []common.Address
var RouterNames map[common.Address]string
var RouterFees map[common.Address]*big.Int

// The token variables.

var TargetTokens []common.Address
var TokenNames map[common.Address]string
var TokenFees map[common.Address]*big.Int

// The network variables.

var EthClient *ethclient.Client
var RpcClient *rpc.Client
var ChainId *big.Int
