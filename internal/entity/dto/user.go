package dto

import (
	"time"
)

type User struct {
	Nickname string    `pq:"nickname"`
	Birthday time.Time `pq:"birthday"`
}
