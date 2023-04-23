package postgres

import "github.com/vladjong/go_project_template/pkg/db/postgres_db"

type repository struct {
	db *postgres_db.PostgresDb
}

func New(db *postgres_db.PostgresDb) *repository {
	return &repository{
		db: db,
	}
}
