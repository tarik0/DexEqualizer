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

type OnSortCallback func(*big.Int, []*circle.TradeOption, []*big.Int, time.Duration, *PairUpdater)

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

	// Action channels.
	tradeCh chan TradeAction
	syncCh  chan SyncAction
	sortCh  chan SortAction

	// Channels.
	blocksCh  chan *types.Header
	pendingCh chan *common.Hash
	logsCh    chan types.Log

	// History channels.
	txHistoryReset  chan bool
	txHistoryAdd    chan HistoryAddAction
	txHistorySearch chan HistorySearchAction

	// Subscriptions
	pendingSub ethereum.Subscription
	blocksSub  ethereum.Subscription
	logsSub    ethereum.Subscription

	// Our transaction history.
	hashToOptionHistory map[common.Hash]*circle.TradeOption
	hashToTxHistory     map[common.Hash]*types.Transaction
	hashToTxBlock       map[common.Hash]*big.Int

	// Pair to minimum gas required to frontrun others.
	pairToMinGasPriceMutex sync.RWMutex
	pairToMinGasPrice      map[common.Address]*big.Int

	// Pending history.
	accountToPendingTx sync.Map

	// Atomic variables.
	syncBlockNum    atomic.Value
	highestBlockNum atomic.Value

	// Other variables.
	params     *PairUpdaterParams
	backend    *ethclient.Client
	rpcBackend *rpc.Client

	pendingBackend *rpc.Client
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
	Path        []common.Address
	Symbols     []string
	Route       []common.Address
	RouteFees   []*big.Int
	RouteTokens [][]common.Address
}

// DebugTraceCall
// A response struct.
type DebugTraceCall struct {
	Calls   []DebugTraceCall `json:"calls"`
	From    common.Address   `json:"from"`
	Gas     string           `json:"gas"`
	GasUsed string           `json:"gasUsed"`
	Input   string           `json:"input"`
	Output  string           `json:"output"`
	To      common.Address   `json:"to"`
	Type    string           `json:"type"`
	Value   string           `json:"value"`
}

// TradeAction
//	Trade action order.
type TradeAction struct {
	BlockNumber         *big.Int
	Transaction         *types.Transaction
	TradeOption         *circle.TradeOption
	ReplacedTransaction *types.Transaction
}

// SyncAction
//	Update reserve order.
type SyncAction struct {
	BlockNumber *big.Int
	TxIndex     *big.Int
	LogIndex    *big.Int
	Address     common.Address
	Res0        *big.Int
	Res1        *big.Int
}

// SortAction
//	Sort circles order.
type SortAction struct {
	BlockNumber *big.Int
}

// HistoryAddAction
//	Add pending transaction to history order.
type HistoryAddAction struct {
	Tx          *types.Transaction
	Option      *circle.TradeOption
	BlockNumber *big.Int
}

// HistorySearchAction
//	Search transaction in the history order.
type HistorySearchAction struct {
	TargetTx       *types.Transaction
	TargetPairAddr common.Address
}

// NewPairUpdater
//	Generates a new pair updater.
func NewPairUpdater(params *PairUpdaterParams, backend *ethclient.Client, rpcBackend *rpc.Client, pendingBackend *rpc.Client) *PairUpdater {
	return &PairUpdater{
		params:         params,
		backend:        backend,
		rpcBackend:     rpcBackend,
		pendingBackend: pendingBackend,
	}
}
