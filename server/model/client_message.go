package model

import (
	"encoding/json"
	"time"
)

type ClientMessageType int

const (
	CLIENT_CREATE_ROOM = 0
	CLIENT_JOIN_ROOM   = 1
	CLIENT_LEAVE_ROOM  = 2
	CLIENT_ERROR       = 3
)

type ClientMessage struct {
	MessageType ClientMessageType `json:"message_type"`
	SenderID    string            `json:"sender_id"`
	Content     string            `json:"content"`
	Timestamp   time.Time         `json:"timestamp"`
}

func (m *ClientMessage) ToJson() []byte {
	val, _ := json.Marshal(m)
	return val
}
