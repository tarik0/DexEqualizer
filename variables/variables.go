package variables

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tarik0/DexEqualizer/abis"
	"github.com/tarik0/DexEqualizer/hub"
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

// Token
//	A basic structure to hold token information.
type Token struct {
	// Basic info.
	Address common.Address
	Symbol  string

	// Fee infos.
	BuyFee  *big.Int
	SellFee *big.Int

	// Gas infos.
	SwapGas *big.Int
}

var TargetTokenAddresses []common.Address
var TargetTokens map[common.Address]*Token

// The network variables.

var EthClient *ethclient.Client
var RpcClient *rpc.Client
var ChainId *big.Int

// The wallet.

var Wallet *wallet.Wallet

// The contracts.

var SwapExec *abis.SwapExecutorV2

// The websocket.

var Hub *hub.Hub
