package model

type RoomInfo struct {
	ID      string       `json:"id"`
	Clients []ClientInfo `json:"clients"`
}
