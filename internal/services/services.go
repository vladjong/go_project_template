package services

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/repository"
)

type Servicer interface {
	Userer
	Notificationer
}

type Userer interface {
	AddUser(ctx context.Context, user entity.User) error
	Users(ctx context.Context) ([]entity.User, error)
}

type Notificationer interface {
	Notifications(ctx context.Context) ([]entity.Notification, error)
}

type Services struct {
	repo repository.Repository

	Userer
	Notificationer
}

func New(repo repository.Repository, opts ...Option) *Services {
	serivces := &Services{
		repo: repo,
	}

	for _, opt := range opts {
		opt(serivces)
	}
	return serivces
}
