package users

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity/dto"
	"github.com/vladjong/go_project_template/internal/repository/postgres/transaction"
)

func (r *Repository) AddUser(ctx context.Context, item dto.User) error {
	query, args, err := r.db.Builder.
		Insert("users").
		Columns("id", "nickname", "birthday").
		Values(item.Id, item.Nickname, item.Birthday).
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
