package postgres_db

import "time"

type Option func(*PostgresDb)

func DSN(dsn string) Option {
	return func(s *PostgresDb) {
		s.cfg.DSN = dsn
	}
}

func MaxIdeConns(connection int) Option {
	return func(s *PostgresDb) {
		s.cfg.MaxIdleConns = connection
	}
}

func MaxOpenConns(connection int) Option {
	return func(s *PostgresDb) {
		s.cfg.MaxOpenConns = connection
	}
}

func ConnMaxLifetime(lifetime time.Duration) Option {
	return func(s *PostgresDb) {
		s.cfg.ConnMaxLifetime = lifetime
	}
}
