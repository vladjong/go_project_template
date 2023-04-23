package users

import (
	"github.com/vladjong/go_project_template/pkg/db/postgres"
)

type Repository struct {
	db *postgres.Postgres
}

func New(db *postgres.Postgres) *Repository {
	return &Repository{
		db: db,
	}
}
