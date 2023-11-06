package v1

import (
	"context"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/vladjong/go_project_template/internal/entity"
	user_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/user"
)

func (s *Service) Create(
	ctx context.Context,
	req *user_grpc.CreateRequest,
) (*user_grpc.CreateResponse, error) {
	switch {
	case req == nil:
		return nil, status.Error(codes.Internal, "empty request")
	case req.Info == nil:
		return nil, status.Error(codes.Internal, "empty info")
	}

	user := entity.UserInfo{
		Nickname: req.Info.Nikname,
		Age:      uint8(req.Info.Age),
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := s.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user_grpc.CreateResponse{
		User: UserToUserGRPC(result),
	}, nil
}
