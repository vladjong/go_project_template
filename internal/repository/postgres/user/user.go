package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/repository/postgres/transaction"
)

func (r *Repository) User(ctx context.Context, id uuid.UUID) (entity.User, error) {
	query, _, err := r.db.Builder.
		Select("*").
		From("users").
		Where("id = ?").
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("error: %w", err)
	}

	user := entity.User{}

	if err := transaction.
		Model(ctx, r.db.Pool).
		QueryRow(ctx, query, id).
    Scan(&user.ID, &user.Nickname, &user.Age, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return entity.User{}, fmt.Errorf("unable to query users: %w", err)
	}
	return user, nil
}
