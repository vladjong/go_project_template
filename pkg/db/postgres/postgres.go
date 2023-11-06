package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	_defaultMaxPoolSize  = 10
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType

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

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	postgres.Pool = pool

	return postgres, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
