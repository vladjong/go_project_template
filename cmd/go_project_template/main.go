package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"

	"github.com/vladjong/go_project_template/configs"
	v1 "github.com/vladjong/go_project_template/internal/controller/http/v1"
	postgres_repository "github.com/vladjong/go_project_template/internal/repository/postgres"
	"github.com/vladjong/go_project_template/internal/services"
	"github.com/vladjong/go_project_template/pkg/db/postgres"
	"github.com/vladjong/go_project_template/pkg/http_server"
	"github.com/vladjong/go_project_template/pkg/logger"
)

func main() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	logger.New(cfg.Logger.Level)

	postgresDriver, err := postgres.New(
		cfg.Postgres.DSN,
		postgres.PoolSize(cfg.Postgres.PoolSize),
		postgres.PoolTimeout(cfg.Postgres.PoolTimeout),
	)
	if err != nil {
		slog.Error("Failed to init postgres driver", err)
		return
	}
	defer postgresDriver.Close()

	repository := postgres_repository.New(
		postgresDriver,
		postgres_repository.InitNotifications(),
		postgres_repository.InitUsers(),
	)

	services := services.New(
		repository,
		services.InitNotifications(),
		services.InitUsers(),
	)

	mux := chi.NewRouter()

	handler := v1.New(mux, *services)
	handler.Run()

	httpServer := http_server.New(
		mux,
		http_server.Port(cfg.HTTP.Port),
		http_server.ReadTimeout(cfg.HTTP.ReadTimeout),
		http_server.ShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("Signal", "signal", s.String())
	case err = <-httpServer.Notify():
		slog.Error("Signal", "HTTP server notify", err)
	}

	if err := httpServer.Shutdown(); err != nil {
		slog.Error("Signal", "HTTP server shutdown", err)
	}
}
