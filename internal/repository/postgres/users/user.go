package users

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity/dto"
	"github.com/vladjong/go_project_template/internal/repository/postgres/transaction"
)

func (r *Repository) User(ctx context.Context, id string) (dto.User, error) {
	query, _, err := r.db.Builder.
		Select("*").
		From("users").
		Where("id = ?").
		ToSql()
	if err != nil {
		return dto.User{}, fmt.Errorf("error: %w", err)
	}

	result := dto.User{}

	if err := transaction.
		Model(ctx, r.db.Pool).
		QueryRow(ctx, query, id).
		Scan(&result.Id, &result.Nickname, &result.Birthday); err != nil {
		return dto.User{}, fmt.Errorf("unable to query users: %w", err)
	}
	return result, nil
}
