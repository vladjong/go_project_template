package v1

import (
	"context"

	"github.com/google/uuid"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/user"
)

type UserServicer interface {
	Create(ctx context.Context, user entity.UserInfo) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
	User(ctx context.Context, id uuid.UUID) (entity.User, error)
}

type Service struct {
	user.UnimplementedUserServiceServer

	userService UserServicer
}

func New(userService UserServicer) *Service {
	return &Service{
		userService: userService,
	}
}
