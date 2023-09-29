package users

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity/dto"
)

func (r *Repository) AddUser(ctx context.Context, item dto.User) error {
	if _, err := r.db.DB.Model(&item).Insert(); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}
