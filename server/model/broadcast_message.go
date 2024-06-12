package model


const (
    BROADCAST_JOIN_ROOM = 1
    BROADCAST_LEAVE_ROOM = 2
    
)

type BroadcastMessage struct {
    MessageType int `json:"message_type"`
    Targets []string `json:"targets"`
    Content string `json:"Content"`
}