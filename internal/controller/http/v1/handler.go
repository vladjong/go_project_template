package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/vladjong/go_project_template/internal/services"
)

type handler struct {
	mux      *chi.Mux
	services services.Services
}

func New(mux *chi.Mux, services services.Services) *handler {
	return &handler{
		mux:      mux,
		services: services,
	}
}

func (h *handler) Run() {
	h.mux.Use(middleware.Logger)
	h.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	h.mux.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.AddUser)
			r.Get("/", h.Users)
		})
	})
}
