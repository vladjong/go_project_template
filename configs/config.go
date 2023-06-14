package configs

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/vladjong/go_project_template/pkg/db/postgres_db"
)

type Config struct {
	APP      App
	HTTP     Http
	Postgres postgres_db.Config
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Http struct {
	Port string `yaml:"port"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("godotend load: %w", err)
	}

	if err := cleanenv.ReadConfig("configs/config.yml", cfg); err != nil {
		return nil, fmt.Errorf(".yml: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf(".env: %w", err)
	}
	return cfg, nil
}
