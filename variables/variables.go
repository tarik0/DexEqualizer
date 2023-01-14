package variables

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/wallet"
	"math/big"
)

// The flags.

var IsDev bool

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

// The wallet.

var Wallet *wallet.Wallet

// The contracts.

var SwapExec *abis.SwapExecutorV2
