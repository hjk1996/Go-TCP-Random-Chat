package model

import (
	"encoding/json"
	"time"
)

type ClientMessageType int

const (
	CLIENT_CREATE_ROOM_CONFIRM   = 0
	CLIENT_JOIN_ROOM_CONFIRM     = 1
	CLIENT_OPPONENT_JOIN_ROOM    = 2
	CLIENT_LEAVE_ROOM_CONFIRM    = 3
	CLIENT_OPPONENT_LEAVE_ROOM   = 4
	CLIENT_OPPONENT_SEND_MESSAGE = 5
	CLIENT_ERROR                 = 6
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
