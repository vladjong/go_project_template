package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/vladjong/go_project_template/internal/entity"
)

type UserRepositorer interface {
	Create(ctx context.Context, item entity.User) error
	List(ctx context.Context) ([]entity.User, error)
	User(ctx context.Context, id uuid.UUID) (entity.User, error)
}

type Service struct {
	userRepo UserRepositorer
}

func New(userRepo UserRepositorer) *Service {
	return &Service{
		userRepo: userRepo,
	}
}
