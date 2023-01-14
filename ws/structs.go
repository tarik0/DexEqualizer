package ws

// WebsocketReq
//	A struct for websocket request.
type WebsocketReq struct {
	Type string `json:"type"`
	Data interface{}
}

// RankReq
// 	A struct for Rank request.
type RankReq struct {
	Circles     interface{} `json:"Circles"`
	SortTime    int64       `json:"SortTime"`    // in ms
	BlockNumber uint64      `json:"BlockNumber"` // in ms
}

// MessageReq
//	A struct for Message request.
type MessageReq struct {
	Timestamp int64  `json:"Timestamp"`
	Message   string `json:"Message"`
}

// HistoryReq
//	A struct for message histories.
type HistoryReq struct {
	Messages []MessageReq `json:"Messages"`
}
