package postgres_db

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/vladjong/go_project_template/pkg/logger"
)

type PostgresDb struct {
	DB  *sqlx.DB
	cfg Config
	log logger.Interface
}

type Config struct {
	DSN             string        `env:"POSTGRES_DSN"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func NewPgx(logger logger.Interface, config Config) (*PostgresDb, error) {
	postgresDb := &PostgresDb{
		cfg: config,
		log: logger,
	}

	db, err := sqlx.Open("pgx", postgresDb.cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open database connection: %w", err)
	}

	db.SetMaxOpenConns(postgresDb.cfg.MaxOpenConns)
	db.SetMaxIdleConns(postgresDb.cfg.MaxIdleConns)
	db.SetConnMaxLifetime(postgresDb.cfg.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	postgresDb.DB = db
	return postgresDb, nil
}

func (p *PostgresDb) Close() {
	if err := p.DB.Close(); err != nil {
		p.log.Error("Can't close database client: %v", err)
	}
}
