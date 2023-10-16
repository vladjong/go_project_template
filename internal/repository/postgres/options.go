package postgres

import (
	"github.com/vladjong/go_project_template/internal/repository/postgres/notifications"
	"github.com/vladjong/go_project_template/internal/repository/postgres/transaction"
	"github.com/vladjong/go_project_template/internal/repository/postgres/users"
)

type Option func(*Repository)

func InitUsers() Option {
	return func(r *Repository) {
		r.Userer = users.New(r.db)
	}
}

func InitNotifications() Option {
	return func(r *Repository) {
		r.Notificationer = notifications.New(r.db)
	}
}

func InitTransactuions() Option {
	return func(r *Repository) {
		r.Transactioner = transaction.New(r.db)
	}
}
