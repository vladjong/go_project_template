package dto

import (
	"time"
)

type User struct {
	Id       string    `db:"id"`
	Nickname string    `db:"nickname"`
	Birthday time.Time `db:"birthday"`
}
