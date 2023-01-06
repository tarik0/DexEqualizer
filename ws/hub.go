package ws

import (
	"github.com/tarik0/DexEqualizer/updater"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// The updater.
	updater *updater.PairUpdater

	// The history.
	history      []MessageReq
	historyMutex *sync.RWMutex
}

// NewHub
//	Generates a new hub.
func NewHub(u *updater.PairUpdater) *Hub {
	return &Hub{
		Broadcast:    make(chan []byte),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
		updater:      u,
		history:      make([]MessageReq, 0),
		historyMutex: &sync.RWMutex{},
	}
}
