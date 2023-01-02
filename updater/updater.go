package updater

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/dexpair"
	"math/big"
	"sync"
	"time"
)

// Config variables.

var MaxProcessAmount = 500
var MaxHops = 3
var MaxCircleResults = 2000

type SyncCallback func(updateTime time.Duration, sortTime time.Duration, u *PairUpdater)

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
	OnSync SyncCallback

	// Other variables.
	params       *PairUpdaterParams
	backend      *ethclient.Client
	sortedTrades []*circle.TradeOption
	logsCh       chan types.Log
	logsSub      ethereum.Subscription
	sortMutex    *sync.RWMutex
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
}

// NewPairUpdater
//	Generates a new pair updater.
func NewPairUpdater(params *PairUpdaterParams, backend *ethclient.Client) *PairUpdater {
	return &PairUpdater{
		params:  params,
		backend: backend,

		sortMutex: &sync.RWMutex{},
	}
}
