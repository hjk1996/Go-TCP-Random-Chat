package server

import "net"

type Room struct {
	ID      string
	Clients map[net.Addr]*Client
}
