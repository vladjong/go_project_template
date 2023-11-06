package grpc

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	v1 "github.com/vladjong/go_project_template/internal/controller/grpc/user/v1"
	"github.com/vladjong/go_project_template/internal/entity"
	user_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/user"
)

type UserServicer interface {
	Create(ctx context.Context, user entity.UserInfo) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
	User(ctx context.Context, id uuid.UUID) (entity.User, error)
}

func RegisterServices(server *grpc.Server, userService UserServicer) {
	user_grpc.RegisterUserServiceServer(server, v1.New(userService))
}
