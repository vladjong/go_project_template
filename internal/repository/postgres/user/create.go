package user

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/repository/postgres/transaction"
)

func (r *Repository) Create(ctx context.Context, item entity.User) error {
	query, args, err := r.db.Builder.
		Insert("users").
		Columns("id", "nickname", "age", "created_at").
		Values(item.ID, item.Nickname, item.Age, item.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := transaction.
		Model(ctx, r.db.Pool).
		Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
