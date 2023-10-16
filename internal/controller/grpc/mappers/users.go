package mappers

import (
	"github.com/vladjong/go_project_template/internal/entity"
	users_grpc "github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/users"
)

func UserToUserGRPC(in entity.User) *users_grpc.User {
	return &users_grpc.User{
		Info: &users_grpc.UserInfo{
			Nikname: in.Nickname,
			// Birthday: imin.Birthday,
		},
	}
}
