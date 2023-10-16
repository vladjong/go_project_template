package notifications

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity/dto"
)

func (r *Repository) CreateNotification(ctx context.Context, notification dto.Notification) (dto.Notification, error) {
	_, err := r.db.Bun.NewInsert().Model(&notification).Exec(ctx)
	if err != nil {
		return dto.Notification{}, err
	}
	return notification, nil
}
