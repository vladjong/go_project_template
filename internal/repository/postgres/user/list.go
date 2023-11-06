package user

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/repository/postgres/transaction"
)

func (r *Repository) List(ctx context.Context) ([]entity.User, error) {
	query, _, err := r.db.Builder.
		Select("*").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	result := []entity.User{}

	rows, err := transaction.Model(ctx, r.db.Pool).Query(ctx, query)
  if err != nil {
    return nil, err
  }

  for rows.Next() {
    user := entity.User{}
    if err := rows.Scan(&user.ID, &user.Nickname, &user.Age, &user.CreatedAt, &user.UpdatedAt); err != nil {
      return nil, err
    }

    result = append(result, user)
  }

	return result, nil
}
