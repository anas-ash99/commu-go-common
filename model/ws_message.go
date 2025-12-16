package model

type WsMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type OnlineMessage struct {
	UserId string `json:"userId"`
	Time   string `json:"time"`
}
