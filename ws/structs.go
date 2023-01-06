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
	Circles    interface{} `json:"Circles"`
	SortTime   int64       `json:"SortTime"`   // in ms
	UpdateTime int64       `json:"UpdateTime"` // in ms
}

// MessageReq
//	A struct for Message request.
type MessageReq struct {
	Message string `json:"Message"`
}

// HistoryReq
//	A struct for message histories.
type HistoryReq struct {
	Messages []MessageReq `json:"Messages"`
}
