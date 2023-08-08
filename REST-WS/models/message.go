package models

type WebSocketMessage struct {
	Type    string      `json:"type"`
	PayLoad interface{} `json:"payload"`
}
