package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"

	"github.com/vladjong/go_project_template/configs"
	v1 "github.com/vladjong/go_project_template/internal/controller/http/v1"
	"github.com/vladjong/go_project_template/internal/repository/postgres"
	"github.com/vladjong/go_project_template/internal/service/peer"
	"github.com/vladjong/go_project_template/pkg/db/postgres_db"
	"github.com/vladjong/go_project_template/pkg/http_server"
	"github.com/vladjong/go_project_template/pkg/logger"
)

func main() {
	log := logger.New("info")

	cfg, err := configs.New()
	if err != nil {
		log.Fatal("Config error: %s", err)
	}
	log.Info("Completed read configs")

	postgresDb, err := postgres_db.NewPgx(log, cfg.Postgres)
	if err != nil {
		log.Fatal("PostgresDb error: %s", err)
	}
	log.Info("Completed init poswgresDb")

	defer postgresDb.Close()

	postgresRepo := postgres.New(postgresDb)

	peerService := peer.New(postgresRepo)

	mux := chi.NewRouter()

	handler := v1.New(mux, peerService, log)
	handler.Run()

	httpServer := http_server.New(mux, log, http_server.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("Signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("HttpServer.Notify: %w", err))
	}

	if err := httpServer.Shutdown(); err != nil {
		log.Error(fmt.Errorf("HttpServer.Shutdown: %w", err))
	}
}
