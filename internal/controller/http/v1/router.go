package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ironstar-io/chizerolog"

	"github.com/vladjong/go_project_template/internal/service/peer"
	"github.com/vladjong/go_project_template/pkg/logger"
)

type handler struct {
	mux         *chi.Mux
	peerService peer.Service
	logger      *logger.Logger
}

func New(mux *chi.Mux, peerService peer.Service, logger *logger.Logger) *handler {
	return &handler{
		mux:         mux,
		peerService: peerService,
		logger:      logger,
	}
}

func (h *handler) Run() {
	h.mux.Use(chizerolog.LoggerMiddleware(h.logger.Logger))
	h.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	h.mux.Route("/api/v1", func(r chi.Router) {
		r.Route("/peer", func(r chi.Router) {
			r.Post("/", h.AddPeer)
			r.Get("/", h.Peers)
		})
	})
}
