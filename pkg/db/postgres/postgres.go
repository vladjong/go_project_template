package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
)

const (
	_defaultPoolSize    = 1
	_defaultPoolTimeout = time.Second
)

type Postgres struct {
	DB *pg.DB

	poolSize    int
	poolTimeout time.Duration
}

type Config struct {
	Name string
}

func New(url string, opts ...Option) (*Postgres, error) {
	postgres := &Postgres{
		poolSize:    _defaultPoolSize,
		poolTimeout: _defaultPoolTimeout,
	}

	for _, opt := range opts {
		opt(postgres)
	}

	opt, err := pg.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("db parse url: %w", err)
	}

	opt.PoolSize = postgres.poolSize
	opt.PoolTimeout = postgres.poolTimeout

	postgres.DB = pg.Connect(opt)

	if err := postgres.DB.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}

	return postgres, nil
}

func (p *Postgres) Close() error {
	if err := p.DB.Close(); err != nil {
		return fmt.Errorf("db close: %w", err)
	}

	return nil
}
