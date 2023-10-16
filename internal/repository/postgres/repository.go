package postgres

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity/dto"
	"github.com/vladjong/go_project_template/pkg/db/postgres"
)

type Userer interface {
	AddUser(ctx context.Context, item dto.User) error
	Users(ctx context.Context) ([]dto.User, error)
	User(ctx context.Context, id string) (dto.User, error)
}

type Notificationer interface {
	Notifications(ctx context.Context) ([]dto.Notification, error)
	CreateNotification(ctx context.Context, notification dto.Notification) (dto.Notification, error)
}

type Transactioner interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
}

type Repository struct {
	db *postgres.Postgres
	Userer
	Notificationer
	Transactioner
}

func New(db *postgres.Postgres, opts ...Option) *Repository {
	repository := &Repository{
		db: db,
	}

	for _, opt := range opts {
		opt(repository)
	}
	return repository
}
