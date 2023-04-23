package entity

import "time"

type Peer struct {
	Nickname string    `json:"nickname"`
	Birthday time.Time `json:"birthday"`
}
