package dto

import "time"

type Peer struct {
	Nickname string    `json:"nickname" db:"nickname"`
	Birthday time.Time `json:"birthday" db:"birthday"`
}
