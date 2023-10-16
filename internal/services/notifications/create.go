package notifications

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/entity/dto"
)

func (s *service) CreateNotification(ctx context.Context, notification entity.Notification) (entity.Notification, error) {
	in := dto.Notification{
		Name: notification.Name,
	}
	out, err := s.repo.CreateNotification(ctx, in)
	if err != nil {
		return entity.Notification{}, err
	}
	return entity.Notification{
		Name: out.Name,
	}, nil
}
