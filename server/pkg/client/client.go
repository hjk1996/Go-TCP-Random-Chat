package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	model "example.com/chat/data"
	server "example.com/chat/internal/server"
)

type ClientState int

const (
	CLIENT_STATE_NO_SERVER_CONNECTION        = 0
	CLIENT_STATE_CONNECTED_BUT_NOT_IN_ROOM   = 1
	CLIENT_STATE_CONNECTED_AND_ROOM_ALONE    = 2
	CLIENT_STATE_CONNECTED_AND_ROOM_TOGETHER = 3
)

type Client struct {
	Address     string
	Conn        net.Conn
	ComChan     chan server.Command
	ClientState ClientState
}

func NewClient(addr string) *Client {

	return &Client{
		Address:     addr,
		ComChan:     make(chan server.Command),
		ClientState: CLIENT_STATE_NO_SERVER_CONNECTION,
	}

}

func (c *Client) Connect() error {

	conn, err := net.Dial("tcp", c.Address)
	defer conn.Close()

	if err != nil {
		return fmt.Errorf("failed to conenect to the server %s : %s", c.Address, err.Error())
	}

	c.Conn = conn
	c.ClientState = CLIENT_STATE_CONNECTED_AND_ROOM_ALONE

	go c.readMessage()
	go c.readInput()

	return nil
}

func (c *Client) readInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		c.printInstruction()
		userInput, _ := reader.ReadString('\n')

	}
}

func (c *Client) printInstruction() {
	switch c.ClientState {
	case CLIENT_STATE_NO_SERVER_CONNECTION:
	case CLIENT_STATE_CONNECTED_BUT_NOT_IN_ROOM:
		print("Press any key to join a room")
	case CLIENT_STATE_CONNECTED_AND_ROOM_ALONE:
		print("Waiting for other user...")
	case CLIENT_STATE_CONNECTED_AND_ROOM_TOGETHER:
		print("message> ")
	}
}

func (c *Client) handleUserInput()

func (c *Client) readMessage() {
	for {
		var msg model.UserMessage

		data, err := bufio.NewReader(c.Conn).ReadBytes('\n')

		if err != nil {
			log.Printf("failed to read message : %s\n", err.Error())
		}

		err = json.Unmarshal(data, &msg)

		if err != nil {
			log.Printf("failed to read message : %s\n", err.Error())
		}

		switch msg.MessageType {
		case model.USER_CREATE_ROOM_CONFIRM:
			c.handleUserCreateRoomConfirm(msg)
		case model.USER_JOIN_ROOM_CONFIRM:
			c.handleUserJoinRoomConfirm(msg)
		case model.OPPONENT_JOIN_ROOM:
			c.handleOpponentJoinRoom(msg)
		case model.USER_LEAVE_ROOM_CONFIRM:
			c.handleUserLeaveRoomConfirm(msg)
		case model.OPPONENT_LEAVE_ROOM:
			c.handleOpponentLeaveRoom(msg)
		case model.OPPONENT_SEND_MESSAGE:
			c.handleOpponentSendMessage(msg)
		case model.ERROR:
			c.handleError(msg)
		}

	}

}

func (c *Client) handleUserCreateRoomConfirm(msg model.UserMessage) {

}

func (c *Client) handleUserJoinRoomConfirm(msg model.UserMessage) {

}

func (c *Client) handleOpponentJoinRoom(msg model.UserMessage) {

}

func (c *Client) handleUserLeaveRoomConfirm(msg model.UserMessage) {

}

func (c *Client) handleOpponentLeaveRoom(msg model.UserMessage) {

}

func (c *Client) handleOpponentSendMessage(msg model.UserMessage) {

}

func (c *Client) handleError(msg model.UserMessage) {

}
