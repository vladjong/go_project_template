package user

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (s *Service) Create(ctx context.Context, userInfo entity.UserInfo) (entity.User, error) {
	user := entity.User{
		ID: uuid.New(),
		UserInfo: entity.UserInfo{
			Nickname: userInfo.Nickname,
			Age:      userInfo.Age,
		},
		CreatedAt: time.Now(),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return entity.User{}, err
	}
	return user, nil
}
