package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/vladjong/go_project_template/configs"
	"github.com/vladjong/go_project_template/internal/controller/grpc"
	user_repo "github.com/vladjong/go_project_template/internal/repository/postgres/user"
	user_serv "github.com/vladjong/go_project_template/internal/service/user"
	"github.com/vladjong/go_project_template/pkg/db/postgres"
	"github.com/vladjong/go_project_template/pkg/grpc_server"
	"github.com/vladjong/go_project_template/pkg/logger"
)

func main() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	logger.New(cfg.Logger.Level)

	postgresDriver, err := postgres.New(
		context.Background(),
		cfg.Postgres.DSN,
		postgres.MaxPoolSize(cfg.Postgres.MaxPoolSize),
		postgres.ConnAttemp(cfg.Postgres.ConnectAttemp),
		postgres.ConnTimeout(cfg.Postgres.ConnectTimeout),
	)
	if err != nil {
		slog.Error("Failed to init postgres driver", err)
		return
	}
	defer postgresDriver.Close()

	userRepo := user_repo.New(postgresDriver)
	userServ := user_serv.New(userRepo)

	grpcServer := grpc_server.New(grpc_server.Port(cfg.GRPC.Port))

	grpc.RegisterServices(grpcServer.Server(), userServ)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Info("Signal", "signal", s.String())
	case err = <-grpcServer.Notify():
		slog.Error("Signal", "GRPC server notify", err)
	}
}
