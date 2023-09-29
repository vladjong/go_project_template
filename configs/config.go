package configs

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	APP      App
	HTTP     Http
	Postgres Postgres
	Logger   Logger
}

type App struct {
	Name    string `yaml:"name" env:"APP_NAME"`
	Version string `yaml:"version" env:"APP_VERSION"`
}

type Http struct {
	Port            string        `yaml:"port" env:"HTTP_PORT"`
	ReadTimeout     time.Duration `yaml:"read_timeout" env:"HTTP_READ_TIMEOUT"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"HTTP_SHUTDOWN_TIMEOUT"`
}

type Postgres struct {
	DSN         string        `yaml:"dsn" env:"POSTGRES_DSN"`
	PoolSize    int           `yaml:"pool_size" env:"POSTGRES_POOL_SIZE"`
	PoolTimeout time.Duration `yaml:"pool_timeout" env:"POSTGRES_POOL_TIMEOUT"`
}

type Logger struct {
	Level string `yaml:"level" env:"LOGGER_LEVEL" env-default:"info"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("configs/config.yml", cfg); err != nil {
		return nil, fmt.Errorf(".yml: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf(".env: %w", err)
	}
	return cfg, nil
}
