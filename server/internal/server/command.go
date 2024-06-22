package server

const (
	CMD_CHANGE_NICK   = 1
	CMD_NEW_ROOM      = 2
	CMD_JOIN_ROOM     = 3
	CMD_LEAVE_ROOM    = 4
	CMD_SEND_MESSAGE  = 5
	CMD_REMOVE_CLIENT = 6
	CMD_QUIT          = 7
)

type Command struct {
	Client        *Client
	CommandType   int
	Args          []string
	CurrentRoomId string
}
