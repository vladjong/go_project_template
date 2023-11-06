package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserInfo

	ID        uuid.UUID `json:"uuid"       validate:"required, uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInfo struct {
	Nickname string `json:"nickname" validate:"required"`
	Age      uint8  `json:"age"      validate:"required"`
}
