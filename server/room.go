package main

import "net"

type Room struct {
	Name    string
	Clients map[net.Addr]*Client
}
