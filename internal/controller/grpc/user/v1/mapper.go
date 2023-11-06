package v1

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/pkg/go-project-template/proto/v1/user"
)

func UsersToUsersGRPC(u []entity.User) []*user.User {
	result := make([]*user.User, len(u))
	for i, v := range u {
		result[i] = UserToUserGRPC(v)
	}
	return result
}

func UserToUserGRPC(u entity.User) *user.User {
	return &user.User{
		Id:        u.ID.String(),
		Info:      UserInfoToUserInfoGRPC(u.UserInfo),
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

func UserInfoToUserInfoGRPC(u entity.UserInfo) *user.UserInfo {
	return &user.UserInfo{
		Nikname: u.Nickname,
		Age:     uint32(u.Age),
	}
}
