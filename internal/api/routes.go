package api

import "net/http"

func NewRoutes(h *Handler) http.Handler {
	mux := http.NewServeMux()

	// Objectives routes
	mux.HandleFunc("/objectives", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetObjectives(w, r)
			return
		}
		if r.Method == http.MethodPost {
			h.CreateObjective(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return mux
}
// Tasks routes
mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetTasks(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.CreateTask(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
})