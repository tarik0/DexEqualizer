package ws

import (
	"github.com/gorilla/websocket"
	"github.com/tarik0/DexEqualizer/updater"
	"net/http"
	"sync"
)

// RankWebsocket
//	A websocket that ranks the circles.
type RankWebsocket struct {
	// The connections.
	connections       []*websocket.Conn
	connectionMutexes map[*websocket.Conn]*sync.RWMutex

	// Other variables.
	upgrader websocket.Upgrader
	updater  *updater.PairUpdater
}

// NewRankWebsocket
// 	Generate new websocket server.
func NewRankWebsocket(u *updater.PairUpdater) *RankWebsocket {
	return &RankWebsocket{
		connectionMutexes: make(map[*websocket.Conn]*sync.RWMutex),
		connections:       make([]*websocket.Conn, 0),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		updater: u,
	}
}
