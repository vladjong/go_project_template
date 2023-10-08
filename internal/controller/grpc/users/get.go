package users

import (
	"context"

	users_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/users"
)

func (s *Service) Get(ctx context.Context, req *users_grpc.GetRequest) (*users_grpc.GetResponse, error) {
	return nil, nil
}
