package users

import (
	"context"

	"github.com/vladjong/go_project_template/internal/controller/grpc/mappers"
	users_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Get(ctx context.Context, req *users_grpc.GetRequest) (*users_grpc.GetResponse, error) {
	// TODO: Validate ID
	user, err := s.service.User(ctx, req.Id)
	if err != nil {
		// TODO: Check Error
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &users_grpc.GetResponse{
		User: mappers.UserToUserGRPC(user),
	}, nil
}
