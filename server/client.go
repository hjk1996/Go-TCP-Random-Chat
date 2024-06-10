package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type Client struct {
	Conn    net.Conn
	ComChan chan<- Command
}

func NewClient(conn net.Conn, comChan chan<- Command) *Client {
	return &Client{
		Conn:    conn,
		ComChan: comChan,
	}

}

func (c *Client) readInput() {
	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {
			log.Printf("Failed to read user input. : %s", err.Error())
		}

		args := strings.Split(msg, " ")

		if len(args) < 2 {
			log.Printf("Invalid input")
			continue
		}

		cmd := args[0]

		switch cmd {
		case "/join":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_JOIN_ROOM,
				Args:        args[1:],
			}

		case "/leave":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_LEAVE_ROOM,
				Args:        args[1:],
			}
		case "/new-room":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_NEW_ROOM,
				Args:        args[1:],
			}
		case "/msg":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_SEND_MESSAGE,
				Args:        args[1:],
			}
		}

	}
}
