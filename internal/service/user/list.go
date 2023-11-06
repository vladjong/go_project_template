package user

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (s *Service) List(ctx context.Context) ([]entity.User, error) {
	users, err := s.userRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("get peers: %w", err)
	}
	return users, nil
}
