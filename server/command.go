package main

const (
    CMD_CHANGE_NICK = 1
    CMD_NEW_ROOM = 2
    CMD_JOIN_ROOM = 3
    CMD_LEAVE_ROOM = 4
    CMD_SEND_MESSAGE = 5
)



type Command struct {
	Client      *Client
	CommandType int
	Args        []string
}
