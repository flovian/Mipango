package api

import "net/http"

func NewRoutes(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.HandleHome)
	mux.HandleFunc("/objectives", h.HandleObjectives)
	mux.HandleFunc("/tasks", h.HandleTasks)
	mux.HandleFunc("/tasks/complete", h.HandleTaskComplete)

	return mux
}
