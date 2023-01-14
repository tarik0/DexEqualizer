package ws

import (
	"bytes"
	"encoding/json"
	"github.com/tarik0/DexEqualizer/circle"
	"github.com/tarik0/DexEqualizer/logger"
	"net/http"
	"time"
)

// Run
//	Starts the hub.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// SetHandler
//	Sets handler for the websocket.
func (h *Hub) SetHandler() {
	http.HandleFunc("/dex_eq", func(w http.ResponseWriter, r *http.Request) {
		client := serveWs(h, w, r)
		h.sendHello(client)
	})
}

// sendHello
//	Sends `History` and `Rank` packets to the new client.
func (h *Hub) sendHello(newClient *Client) {
	// Get trade options.
	options := h.updater.GetSortedTrades()
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

	// Encoder.
	var buff = new(bytes.Buffer)
	e := json.NewEncoder(buff)
	e.SetEscapeHTML(true)

	// Marshall history.
	err := e.Encode(WebsocketReq{
		Type: "History",
		Data: HistoryReq{
			Messages: h.history,
		},
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to marshal history.")
	}

	// Broadcast
	newClient.send <- buff.Bytes()

	// Get block numbers.
	blockNum := h.updater.GetBlockNumber()

	// Empty buffer.
	buff = new(bytes.Buffer)
	e = json.NewEncoder(buff)
	e.SetEscapeHTML(true)

	// Marshall rank.
	err = e.Encode(WebsocketReq{
		Type: "Rank",
		Data: RankReq{
			Circles:     tradesJson,
			SortTime:    h.updater.GetSortTime().Milliseconds(),
			BlockNumber: blockNum,
		},
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to marshal rank.")
	}

	// Broadcast
	newClient.send <- buff.Bytes()
}

// AddToHistory
//	Adds message to the history.
func (h *Hub) AddToHistory(str MessageReq) {
	logger.Log.Infoln(str.Message)

	h.historyMutex.Lock()
	h.history = append(h.history, str)
	h.historyMutex.Unlock()
}

// ClearHistory
//	A goroutine to clear the history periodically.
func (h *Hub) ClearHistory() {
	for {
		time.Sleep(historyCleanInterval)

		h.historyMutex.Lock()
		h.history = make([]MessageReq, 0)
		h.historyMutex.Unlock()
	}
}
