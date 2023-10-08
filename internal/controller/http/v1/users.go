package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/metrics"
)

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	start := time.Now()
	statusCode := http.StatusOK
	defer func() {
		metrics.ObserveRequest(time.Since(start), statusCode)
	}()

	peer := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		statusCode = http.StatusBadRequest
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.services.AddUser(ctx, peer); err != nil {
		statusCode = http.StatusInternalServerError
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
}

func (h *handler) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	start := time.Now()
	statusCode := http.StatusOK
	defer func() {
		metrics.ObserveRequest(time.Since(start), statusCode)
	}()

	ctx := context.Background()

	peers, err := h.services.Users(ctx)
	if err != nil {
		statusCode := http.StatusInternalServerError
		WriteError(w, err, statusCode)
		return
	}

	WriteResponseJson(w, peers)
}
