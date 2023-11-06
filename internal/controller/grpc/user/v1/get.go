package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	user_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/user"
)

func (s *Service) Get(
	ctx context.Context,
	req *user_grpc.GetRequest,
) (*user_grpc.GetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Internal, "empty request")
	}

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("id (%s)", req.Id))
	}

	result, err := s.userService.User(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user_grpc.GetResponse{
		User: UserToUserGRPC(result),
	}, nil
}
