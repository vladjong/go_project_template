package users

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity/dto"
)

func (r *Repository) User(ctx context.Context, id string) (dto.User, error) {
	return dto.User{}, nil
}
