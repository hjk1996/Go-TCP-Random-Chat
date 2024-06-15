package model

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type RoomInfo struct {
	ID      string       `json:"id"`
	Clients []ClientInfo `json:"clients"`
}

func (r *RoomInfo) Copy() *RoomInfo {
	var b bytes.Buffer
	var result RoomInfo
	e := gob.NewEncoder(&b)
	d := gob.NewDecoder(&b)

	e.Encode(r)

	d.Decode(&result)

	return &result

}

func (r *RoomInfo) ToJson() []byte {
	val, _ := json.Marshal(r)
	return val
}
