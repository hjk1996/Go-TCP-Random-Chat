package model

import "encoding/json"

type BroadcastMessageType int

const (
	BROADCAST_JOIN_ROOM    = 1
	BROADCAST_LEAVE_ROOM   = 2
	BROADCAST_SEND_MESSAGE = 3
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
