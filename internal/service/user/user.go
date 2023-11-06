package user

import (
	"context"

	"github.com/google/uuid"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (s *Service) User(ctx context.Context, id uuid.UUID) (entity.User, error) {
	user, err := s.userRepo.User(ctx, id)
	if err != nil {
		return entity.User{}, nil
	}

	return user, nil
}
