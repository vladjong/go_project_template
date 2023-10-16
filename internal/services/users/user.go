package users

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/entity/dto"
)

func (s *service) User(ctx context.Context, id string) (entity.User, error) {
	var (
		err  error
		user dto.User
	)

	if err := s.repo.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.repo.AddUser(txCtx, dto.User{
			Id:       uuid.NewString(),
			Nickname: "test",
			Birthday: time.Now(),
		}); err != nil {
			return err
		}

		user, err = s.repo.User(txCtx, id)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return entity.User{}, err
	}
	return entity.User{
		Nickname: user.Nickname,
		Birthday: user.Birthday,
	}, err
}
