package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/metrics"
)

func (h *handler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	start := time.Now()
	statusCode := http.StatusOK
	defer func() {
		metrics.ObserveRequest(time.Since(start), statusCode)
	}()

	notification := entity.Notification{}
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		statusCode = http.StatusBadRequest
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	result, err := h.services.CreateNotification(ctx, notification)
	if err != nil {
		statusCode = http.StatusInternalServerError
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, result)
}
