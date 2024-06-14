package model

import "encoding/json"

type BroadcastMessageType int

const (
	BROADCAST_JOIN_ROOM    = 0
	BROADCAST_LEAVE_ROOM   = 1
	BROADCAST_SEND_MESSAGE = 2
)

type BroadcastMessage struct {
	MessageType BroadcastMessageType `json:"message_type"`
	SenderId    string               `json:"sender_id"`
	Targets     []string             `json:"targets"`
	Content     string               `json:"Content"`
}

func (m *BroadcastMessage) ToJson() []byte {
	val, _ := json.Marshal(m)
	return val
}
