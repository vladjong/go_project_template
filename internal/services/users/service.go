package users

import (
	"github.com/vladjong/go_project_template/internal/repository"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}
