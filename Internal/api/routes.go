package api

import (
	"net/http"
)

func NewRoutes(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/objectives", h.GetObjectives)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	return mux
}