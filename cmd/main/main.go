package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"mipango/internal/repository"
	"mipango/internal/services"
)

func main() {
	objRepo := repository.NewObjectiveRepo()
	taskRepo := services.NewTaskRepo()
	service := services.NewObjectiveService(objRepo, taskRepo)

	tmplIndex := template.Must(template.ParseFiles("templates/index.html"))
	tmplObjectives := template.Must(template.ParseFiles("templates/objectives.html"))
	tmplTasks := template.Must(template.ParseFiles("templates/tasks.html"))

	// Home
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplIndex.Execute(w, nil)
	})

	// Objectives
	http.HandleFunc("/objectives", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmplObjectives.Execute(w, service.GetAllObjectives())
			return
		}
		if r.Method == http.MethodPost {
			title := r.FormValue("title")
			deadline := r.FormValue("deadline")
			priority, _ := strconv.Atoi(r.FormValue("priority"))
			service.CreateObjective(title, deadline, priority)
			http.Redirect(w, r, "/objectives", http.StatusSeeOther)
		}
	})

	// Tasks
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		objID := r.FormValue("objective_id")
		if r.Method == http.MethodGet {
			// Find objective
			var obj interface{}
			for _, o := range service.GetAllObjectives() {
				if o.ID == objID {
					obj = o
				}
			}
			data := struct {
				Objective interface{}
				Tasks     []*services.TaskRepo
			}{
				Objective: obj,
				Tasks:     service.GetTasks(objID),
			}
			tmplTasks.Execute(w, data)
			return
		}
		if r.Method == http.MethodPost {
			title := r.FormValue("title")
			deadline := r.FormValue("deadline")
			priority, _ := strconv.Atoi(r.FormValue("priority"))
			service.CreateTask(title, objID, priority, deadline)
			http.Redirect(w, r, "/tasks?objective_id="+objID, http.StatusSeeOther)
		}
	})

	port := ":8080"
	log.Println("Mipango server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}