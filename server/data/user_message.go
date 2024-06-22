package model

import (
	"encoding/json"
	"time"
)

type UserMessageType int

const (
	USER_CREATE_ROOM_CONFIRM = 0
	USER_JOIN_ROOM_CONFIRM   = 1
	OPPONENT_JOIN_ROOM       = 2
	USER_LEAVE_ROOM_CONFIRM  = 3
	OPPONENT_LEAVE_ROOM      = 4
	OPPONENT_SEND_MESSAGE    = 5
	ERROR                    = 6
)

type UserMessage struct {
	MessageType UserMessageType `json:"message_type"`
	SenderID    string          `json:"sender_id"`
	Content     string          `json:"content"`
	Timestamp   time.Time       `json:"timestamp"`
}

func (m *UserMessage) ToJson() []byte {
	val, _ := json.Marshal(m)
	return val
}
