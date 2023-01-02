package ws

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/logger"
	"net/http"
	"sync"
)

var addr = flag.String("Web Socket Address", "localhost:8081", "Dex Equalizer WebSocket")

// Start
//	Starts the server.
func (ws *RankWebsocket) Start() error {
	http.HandleFunc("/dex_eq", ws.UpgradeRequest)
	return http.ListenAndServe(*addr, nil)
}

// UpgradeRequest
//	Upgrades the requests and adds writer to the collection.
func (ws *RankWebsocket) UpgradeRequest(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection.
	c, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to upgrade request.")
	}

	// Get trade options.
	options := ws.updater.GetSortedTrades()
	if options == nil {
		return
	}

	// Print the best 10 options.
	var tradesJson = make([]circle.TradeOptionJSON, 5)
	for i, opt := range options {
		tradesJson[i] = opt.GetJSON()
		if i == 4 {
			break
		}
	}

	// Append to connections.
	ws.connectionMutexes[c] = &sync.RWMutex{}
	ws.connections = append(ws.connections, c)

	// Marshall.
	rankBytes, err := json.Marshal(WebsocketReq{
		Type: "Rank",
		Data: RankReq{
			Circles:    tradesJson,
			SortTime:   0,
			UpdateTime: 0,
		},
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to marshal trade.")
	}

	// Broadcast
	ws.Broadcast(rankBytes)
}

// Broadcast
//	Sends infos to all connections.
func (ws *RankWebsocket) Broadcast(str []byte) {
	for i, con := range ws.connections {
		ws.connectionMutexes[con].Lock()
		err := con.WriteMessage(websocket.TextMessage, str)
		ws.connectionMutexes[con].Unlock()
		if err != nil {
			delete(ws.connectionMutexes, con)
			ws.connections = append(ws.connections[:i], ws.connections[i+1:]...)
		}
	}
}
