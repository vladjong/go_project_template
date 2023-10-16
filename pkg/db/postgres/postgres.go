package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	_defaultMaxPoolSize  = 10
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType
	Bun     *bun.DB

	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
}

type Config struct {
	Name string
}

func New(ctx context.Context, url string, opts ...Option) (*Postgres, error) {
	postgres := &Postgres{}

	for _, opt := range opts {
		opt(postgres)
	}

	postgres.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// pool, err := pgxpool.New(ctx, url)
	// if err != nil {
	// 	return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	// }

	// postgres.Pool = pool
	fmt.Println(url)

	postgres.Bun = bun.NewDB(sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(url))),
		pgdialect.New())

	if err := postgres.Bun.Ping(); err != nil {
		return nil, err
	}

	return postgres, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
