package v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	user_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/user"
)

func (s *Service) List(ctx context.Context, req *emptypb.Empty) (*user_grpc.ListResponse, error) {
	result, err := s.userService.List(ctx)
	if err != nil {
		return nil, err
	}

	return &user_grpc.ListResponse{
		Users: UsersToUsersGRPC(result),
	}, nil
}
