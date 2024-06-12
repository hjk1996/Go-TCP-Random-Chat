package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

type Client struct {
	ID            string
	Conn          net.Conn
	ComChan       chan<- Command
	CurrentRoomId string
}

func NewClient(conn net.Conn, comChan chan<- Command) *Client {
	clientId := uuid.New().String()
	return &Client{
		ID:      clientId,
		Conn:    conn,
		ComChan: comChan,
	}
}

func (c *Client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection closed by client %s", c.ID)
				c.ComChan <- Command{
					Client:      c,
					CommandType: CMD_REMOVE_CLIENT,
				}
				return
			}
			log.Printf("Failed to read user input: %s", err.Error())
		}

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/join":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_JOIN_ROOM,
			}

		case "/leave":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_LEAVE_ROOM,
			}
		case "/new-room":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_NEW_ROOM,
			}
		case "/msg":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_SEND_MESSAGE,
				Args:        args,
			}
		default:
			log.Printf("Wrong command: %s", cmd)
		}

	}
}
