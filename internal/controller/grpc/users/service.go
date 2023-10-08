package users

import (
	"github.com/vladjong/go_project_template/internal/services"
	users_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/users"
)

type Service struct {
	users_grpc.UnimplementedUsersServer
	service services.Services
}

func New(service services.Services) *Service {
	return &Service{
		service: service,
	}
}
