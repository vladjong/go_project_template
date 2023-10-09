package users

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (s *service) User(ctx context.Context, id string) (entity.User, error) {
	return entity.User{}, nil
}
