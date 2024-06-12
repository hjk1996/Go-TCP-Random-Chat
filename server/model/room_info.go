package model

import (
	"bytes"
	"encoding/gob"
)

type RoomInfo struct {
	ID      string       `json:"id"`
	Clients []ClientInfo `json:"clients"`
}

func (r *RoomInfo) Copy() (*RoomInfo, error) {
	var b bytes.Buffer
	var result RoomInfo
	e := gob.NewEncoder(&b)
	d := gob.NewDecoder(&b)

	err := e.Encode(r)
	if err != nil {
		return nil, err
	}
	err = d.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
