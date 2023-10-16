package postgres

import "time"

type Option func(*Postgres)

func MaxPoolSize(size int) Option {
	return func(p *Postgres) {
		p.maxPoolSize = size
	}
}

func ConnTimeout(timeout time.Duration) Option {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}

func ConnAttemp(attemp int) Option {
	return func(c *Postgres) {
		c.connAttempts = attemp
	}
}
