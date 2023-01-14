package updater

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/dexpair"
	"math/big"
	"sync/atomic"
)

// Config variables.

var MaxProcessAmount = 750
var MaxHops = 3
var MaxCircleResults = 10000

type OnSortCallback func(header *types.Header, u *PairUpdater)

// PairUpdater
//	A system that checks pair reserves for arbitrage options.
type PairUpdater struct {
	// Token maps.
	TokenToDecimals map[common.Address]*big.Int
	TokenToPairs    map[common.Address][]*dexpair.DexPair
	PairToCircles   map[common.Address]map[uint64]*circle.Circle

	// Pair maps.
	AddressToPair map[common.Address]*dexpair.DexPair
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
	logsCh   chan types.Log
	blocksCh chan *types.Header

	// Subscriptions
	logsSub   ethereum.Subscription
	blocksSub ethereum.Subscription

	// Atomic variables.
	lastBlockNum atomic.Value
	sortedTrades atomic.Value
	sortTime     atomic.Value

	// Other variables.
	params  *PairUpdaterParams
	backend *ethclient.Client
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
	MaxResultCount int
	MaxHops        int
	Path           []common.Address
	Symbols        []string
	Route          []common.Address
	RouteFees      []*big.Int
	RouteTokens    [][]common.Address
	RouteReserves  [][]*big.Int
}

// NewPairUpdater
//	Generates a new pair updater.
func NewPairUpdater(params *PairUpdaterParams, backend *ethclient.Client) *PairUpdater {
	return &PairUpdater{
		params:  params,
		backend: backend,
	}
}
