package users

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (s *service) Users(ctx context.Context) ([]entity.User, error) {
	usersDto, err := s.repo.Users(ctx)
	if err != nil {
		return nil, fmt.Errorf("get peers: %w", err)
	}

	users := make([]entity.User, len(usersDto))
	for i, user := range usersDto {
		users[i] = entity.User{
			Nickname: user.Nickname,
			Birthday: user.Birthday,
		}
	}
	return users, nil
}
