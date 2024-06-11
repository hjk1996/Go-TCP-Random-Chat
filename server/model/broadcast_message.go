package model


const (
    BROADCAST_LEAVE_ROOM = 1
)

type BroadcastMessage struct {
    MessageType int
    Targets []string
    Content string
}