package postgres

import "time"

type Option func(*Postgres)

func PoolSize(size int) Option {
	return func(p *Postgres) {
		p.poolSize = size
	}
}

func PoolTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.poolTimeout = timeout
	}
}
