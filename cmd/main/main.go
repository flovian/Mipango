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
	objService := services.NewObjectiveService(objRepo, taskRepo)

	tmplIndex := template.Must(template.ParseFiles("templates/index.html"))
	tmplObjectives := template.Must(template.ParseFiles("templates/objectives.html"))
	tmplTasks := template.Must(template.ParseFiles("templates/tasks.html"))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplIndex.Execute(w, nil)
	})

	// Objectives page
	http.HandleFunc("/objectives", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			objectives := objService.GetAllObjectives()
			tmplObjectives.Execute(w, objectives)
			return
		}

		if r.Method == http.MethodPost {
			title := r.FormValue("title")
			deadline := r.FormValue("deadline")
			priorityStr := r.FormValue("priority")
			priority, _ := strconv.Atoi(priorityStr)
			objService.CreateObjective(title, deadline, priority)
			http.Redirect(w, r, "/objectives", http.StatusSeeOther)
		}
	})

	// Tasks page
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		objectiveID := r.FormValue("objective_id")
		if r.Method == http.MethodGet {
			objectives := objService.GetAllObjectives()
			var currentObj *services.ObjectiveService
			for _, obj := range objectives {
				if obj.ID == objectiveID {
					currentObj = &services.ObjectiveService{}
				}
			}
			data := struct {
				Objective interface{}
				Tasks     interface{}
			}{
				Objective: objService.GetAllObjectives(), // For simplicity, you can map properly
				Tasks:     objService.GetTasks(objectiveID),
			}
			tmplTasks.Execute(w, data)
			return
		}

		if r.Method == http.MethodPost {
			title := r.FormValue("title")
			deadline := r.FormValue("deadline")
			priorityStr := r.FormValue("priority")
			priority, _ := strconv.Atoi(priorityStr)
			objService.CreateTask(title, objectiveID, priority, deadline)
			http.Redirect(w, r, "/tasks?objective_id="+objectiveID, http.StatusSeeOther)
		}
	})

	port := ":8080"
	log.Println("🚀 Mipango server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}