package api

import (
	"html/template"
	"net/http"
	"strconv"

	"mipango/internal/services"
)

type Handler struct {
	ObjectiveService *services.ObjectiveService
	Templates        *template.Template
}

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	objs, _ := h.ObjectiveService.GetAllObjectives()
	suggestion, _ := h.ObjectiveService.GetSmartSuggestion()
	data := map[string]interface{}{
		"Objectives": objs,
		"Suggestion": suggestion,
	}
	h.Templates.ExecuteTemplate(w, "index.html", data)
}

func (h *Handler) HandleObjectives(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		objs, _ := h.ObjectiveService.GetAllObjectives()
		h.Templates.ExecuteTemplate(w, "objective.html", objs)
		return
	}
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		deadline := r.FormValue("deadline")
		priority, _ := strconv.Atoi(r.FormValue("priority"))
		h.ObjectiveService.CreateObjective(title, deadline, priority)
		http.Redirect(w, r, "/objectives", http.StatusSeeOther)
	}
}

func (h *Handler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	objID := r.URL.Query().Get("objective_id")
	if r.Method == http.MethodGet {
		tasks, _ := h.ObjectiveService.GetTasks(objID)
		objs, _ := h.ObjectiveService.GetAllObjectives()
		var currentObj interface{}
		for _, o := range objs {
			if o.ID == objID {
				currentObj = o
				break
			}
		}
		data := map[string]interface{}{
			"Objective": currentObj,
			"Tasks":     tasks,
		}
		h.Templates.ExecuteTemplate(w, "tasks.html", data)
		return
	}
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		deadline := r.FormValue("deadline")
		priority, _ := strconv.Atoi(r.FormValue("priority"))
		h.ObjectiveService.CreateTask(title, objID, priority, deadline)
		http.Redirect(w, r, "/tasks?objective_id="+objID, http.StatusSeeOther)
	}
}

func (h *Handler) HandleTaskComplete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		taskID := r.FormValue("task_id")
		objID := r.FormValue("objective_id")
		completed := r.FormValue("completed") == "on"
		h.ObjectiveService.UpdateTaskCompletion(taskID, completed)
		http.Redirect(w, r, "/tasks?objective_id="+objID, http.StatusSeeOther)
	}
}
