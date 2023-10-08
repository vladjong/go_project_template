package metrics

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type handler struct {
	mux *chi.Mux
}

func New(mux *chi.Mux) *handler {
	return &handler{
		mux: mux,
	}
}

func (h *handler) Run() {
	h.mux.Use(middleware.Logger)
	h.mux.Handle("/metrics", promhttp.Handler())
}
