package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vladjong/go_project_template/internal/entity"
)

func (h *handler) AddPeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	peer := entity.Peer{}
	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.peerService.AddPeer(ctx, peer); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Peers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	peers, err := h.peerService.GetPeers(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peers)
}
