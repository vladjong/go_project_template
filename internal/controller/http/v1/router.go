package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vladjong/go_project_template/internal/service/peer"
)

type handler struct {
	mux         *chi.Mux
	peerService peer.Service
}

func New(mux *chi.Mux, peerService peer.Service) *handler {
	return &handler{
		mux:         mux,
		peerService: peerService,
	}
}

func (h *handler) Run() {
	h.mux.Use(middleware.Logger)

	h.mux.Post("/api/v1/peer", h.AddPeer)
	h.mux.Get("/api/v1/peer", h.Peers)
}
