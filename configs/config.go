package configs

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	APP      *App
	HTTP     *Http
	Postgres *Postgres
}

type App struct {
	Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type Http struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type Postgres struct {
	DSN             string        `env-required:"true" yaml:"database_dsn" env:"DATABASE_DSN"`
	MaxOpenConns    int           `env-required:"true" yaml:"max_open_conns" env:"DATABASE_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `env-required:"true" yaml:"max_idle_conns" env:"DATABASE_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `env-required:"true" yaml:"conn_max_lifetime" env:"DATABASE_CONN_MAX_LIFETIME"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("configs/config.yaml", cfg); err != nil {
		return nil, fmt.Errorf(".yaml: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf(".env: %w", err)
	}

	return cfg, nil
}
