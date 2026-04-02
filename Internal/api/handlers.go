package api

import (
	"encoding/json"
	"net/http"
	"mipango/internal/services"
)

type Handler struct {
	ObjectiveService *services.ObjectiveService
}

func (h *Handler) GetObjectives(w http.ResponseWriter, r *http.Request) {
	objs := h.ObjectiveService.GetAllObjectives()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}