package users

import (
	"github.com/vladjong/go_project_template/internal/services"
	users_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/users"
)

type Service struct {
	users_grpc.UnimplementedUsersServer
	service services.Servicer
}

func New(service services.Servicer) *Service {
	return &Service{
		service: service,
	}
}
