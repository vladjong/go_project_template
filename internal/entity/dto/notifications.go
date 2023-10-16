package dto

import "github.com/uptrace/bun"

type Notification struct {
	bun.BaseModel `bun:"table:notifications"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Name          string `bun:"name,notnull,type:varchar(64)"`
}
