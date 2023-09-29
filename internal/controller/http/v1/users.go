package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	peer := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.services.Users.AddUser(ctx, peer); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	peers, err := h.services.Users.Users(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peers)
}
