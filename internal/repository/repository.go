package repository

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity/dto"
)

type Repository interface {
	Userer
	Notificationer
	Transactioner
}

type Userer interface {
	AddUser(ctx context.Context, item dto.User) error
	Users(ctx context.Context) ([]dto.User, error)
	User(ctx context.Context, id string) (dto.User, error)
}

type Notificationer interface {
	CreateNotification(ctx context.Context, notification dto.Notification) (dto.Notification, error)
	Notifications(ctx context.Context) ([]dto.Notification, error)
}

type Transactioner interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
}
