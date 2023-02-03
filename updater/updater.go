package updater

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/dexpair"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

// Config variables.

var MaxProcessAmount = 10000

type OnSortCallback func(*types.Header, []*circle.TradeOption, time.Duration, *PairUpdater)

// PairUpdater
//	A system that checks pair reserves for arbitrage options.
type PairUpdater struct {
	// Token maps.
	TokenToDecimals map[common.Address]*big.Int
	TokenToPairs    map[common.Address][]*dexpair.DexPair

	// Pair maps.
	AddressToPair map[common.Address]*dexpair.DexPair
	PairAddresses []common.Address
	PairToTokens  map[common.Address][]common.Address
	PairToFactory map[common.Address]common.Address

	// Router maps.
	RouterToFactory map[common.Address]common.Address

	// Factory maps.
	FactoryToRouter map[common.Address]common.Address

	// Arrays.
	Factories []common.Address
	Pairs     []*dexpair.DexPair
	Circles   map[uint64]*circle.Circle

	// The callbacks.
	OnSort OnSortCallback

	// Channels.
	blocksCh  chan *types.Header
	pendingCh chan *common.Hash

	// History channels.
	TxHistoryReset chan bool
	TxHistoryAdd   chan struct {
		Tx          *types.Transaction
		Option      *circle.TradeOption
		BlockNumber *big.Int
	}
	TxHistorySearch chan struct {
		TargetTx       *types.Transaction
		TargetPairAddr common.Address
	}

	// Subscriptions
	pendingSub ethereum.Subscription
	blocksSub  ethereum.Subscription

	// Our transaction history.
	hashToOptionHistory map[common.Hash]*circle.TradeOption
	hashToTxHistory     map[common.Hash]*types.Transaction
	hashToTxBlock       map[common.Hash]*big.Int

	// Pending history.
	accountToPendingTx sync.Map

	// Atomic variables.
	lastBlockNum    atomic.Value
	filterLogsMutex sync.RWMutex

	// Other variables.
	params     *PairUpdaterParams
	backend    *ethclient.Client
	rpcBackend *rpc.Client
}

// PairUpdaterParams
//	Parameter struct for generating a new PairUpdater.
type PairUpdaterParams struct {
	// Router information.
	Routers struct {
		Addresses []common.Address
		Names     map[common.Address]string
		Fees      map[common.Address]*big.Int
	}

	// Token information.
	Tokens struct {
		MainAddress common.Address
		Addresses   []common.Address
		Symbols     map[common.Address]string
		Fees        map[common.Address]*big.Int
	}

	// Multicaller information.
	Multicaller struct {
		Address common.Address
	}
}

// DFSCircleParams
//	Helper struct to make recursive things easier.
type DFSCircleParams struct {
	Path          []common.Address
	Symbols       []string
	Route         []common.Address
	RouteFees     []*big.Int
	RouteTokens   [][]common.Address
	RouteReserves [][]*big.Int
}

// DebugTraceCall
// A response struct.
type DebugTraceCall struct {
	Failed      bool   `json:"failed"`
	Gas         int    `json:"gas"`
	ReturnValue string `json:"returnValue"`
	StructLogs  []struct {
		Depth   int      `json:"depth"`
		Gas     int      `json:"gas"`
		GasCost int      `json:"gasCost"`
		Op      string   `json:"op"`
		Pc      int      `json:"PC"`
		Stack   []string `json:"stack"`
	} `json:"structLogs"`
}

// NewPairUpdater
//	Generates a new pair updater.
func NewPairUpdater(params *PairUpdaterParams, backend *ethclient.Client, rpcBackend *rpc.Client) *PairUpdater {
	return &PairUpdater{
		params:     params,
		backend:    backend,
		rpcBackend: rpcBackend,
	}
}
