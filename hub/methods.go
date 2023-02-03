package hub

import (
	"bytes"
	"encoding/json"
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
		case message := <-h.broadcast:
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

	// broadcast
	newClient.send <- buff.Bytes()

	// Empty buffer.
	buff = new(bytes.Buffer)
	e = json.NewEncoder(buff)
	e.SetEscapeHTML(true)

	// Marshall rank.
	err = e.Encode(WebsocketReq{
		Type: "Rank",
		Data: RankReq{
			Circles:     nil,
			SortTime:    0,
			BlockNumber: 0,
		},
	})
	if err != nil {
		logger.Log.WithError(err).Fatalln("Unable to marshal rank.")
	}

	// broadcast
	newClient.send <- buff.Bytes()
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

// BroadcastMsg broadcasts a message and adds it to the history.
func (h *Hub) BroadcastMsg(msg string) error {
	// Encoder.
	var buff = new(bytes.Buffer)
	e := json.NewEncoder(buff)
	e.SetEscapeHTML(true)

	// Marshall.
	messageReq := MessageReq{
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}

	// Append to history.
	h.historyMutex.Lock()
	h.history = append(h.history, messageReq)
	h.historyMutex.Unlock()

	// Encode.
	err := e.Encode(WebsocketReq{
		Type: "Message",
		Data: messageReq,
	})
	if err != nil {
		return err
	}

	h.broadcast <- buff.Bytes()
	return nil
}

// BroadcastRanks broadcast a rank message.
func (h *Hub) BroadcastRanks(trades interface{}, sortTime int64, blockNumber uint64) error {
	// Encoder.
	var buff = new(bytes.Buffer)
	e := json.NewEncoder(buff)
	e.SetEscapeHTML(true)

	// Marshall.
	err := e.Encode(WebsocketReq{
		Type: "Rank",
		Data: RankReq{
			Circles:     trades,
			SortTime:    sortTime,
			BlockNumber: blockNumber,
		},
	})
	if err != nil {
		return err
	}

	h.broadcast <- buff.Bytes()
	return nil
}
