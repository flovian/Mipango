package api

import (
	"encoding/json"
	"net/http"

	"mipango/internal/services"
)

type Handler struct {
	ObjectiveService *services.ObjectiveService
}

// GET /objectives
func (h *Handler) GetObjectives(w http.ResponseWriter, r *http.Request) {
	objs := h.ObjectiveService.GetAllObjectives()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

// Request body struct
type CreateObjectiveRequest struct {
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
	Priority int    `json:"priority"`
}

// POST /objectives
func (h *Handler) CreateObjective(w http.ResponseWriter, r *http.Request) {
	var req CreateObjectiveRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	obj := h.ObjectiveService.CreateObjective(req.Title, req.Deadline, req.Priority)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
// ----------------- TASK HANDLERS -----------------

type CreateTaskRequest struct {
	Title       string `json:"title"`
	ObjectiveID string `json:"objective_id"`
	Priority    int    `json:"priority"`
	Deadline    string `json:"deadline"`
}

// POST /tasks
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := h.ObjectiveService.CreateTask(req.Title, req.ObjectiveID, req.Priority, req.Deadline)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GET /tasks?objective_id=<id>
func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	objID := r.URL.Query().Get("objective_id")
	tasks := h.ObjectiveService.GetTasks(objID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}