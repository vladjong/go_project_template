package services

import (
	"github.com/vladjong/go_project_template/internal/services/notifications"
	"github.com/vladjong/go_project_template/internal/services/users"
)

type Option func(*Services)

func InitUsers() Option {
	return func(s *Services) {
		s.Userer = users.New(s.repo)
	}
}

func InitNotifications() Option {
	return func(s *Services) {
		s.Notificationer = notifications.New(s.repo)
	}
}
