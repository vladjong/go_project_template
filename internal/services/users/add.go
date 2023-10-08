package users

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/entity/dto"
)

func (s *service) AddUser(ctx context.Context, user entity.User) error {
	userDto := dto.User{
		Nickname: user.Nickname,
		Birthday: user.Birthday,
	}

	if err := s.repo.AddUser(ctx, userDto); err != nil {
		return fmt.Errorf("add user: %w", err)
	}
	return nil
}
