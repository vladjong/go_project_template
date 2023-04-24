package configs

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	APP      App
	HTTP     Http
	Postgres Postgres
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Http struct {
	Port string `yaml:"port"`
}

type Postgres struct {
	DSN             string        `env:"POSTGRES_DSN"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func New() (*Config, error) {
	cfg := &Config{}
	godotenv.Load()

	if err := cleanenv.ReadConfig("configs/config.yml", cfg); err != nil {
		return nil, fmt.Errorf(".yml: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf(".env: %w", err)
	}
	return cfg, nil
}
