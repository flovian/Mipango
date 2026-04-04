package main

import (
	"html/template"
	"log"
	"net/http"

	"mipango/internal/api"
	"mipango/internal/repository"
	"mipango/internal/services"
)

func main() {
	// Initialize DB
	db := repository.InitDB("mipango.db")
	defer db.Close()

	// Init Repos
	objRepo := repository.NewObjectiveRepo(db)
	taskRepo := repository.NewTaskRepo(db)

	// Init Service
	service := services.NewObjectiveService(objRepo, taskRepo)

	// Parse Templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Init Handlers
	handler := &api.Handler{
		ObjectiveService: service,
		Templates:        tmpl,
	}

	// Serve Static Files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Init Routes
	mux := api.NewRoutes(handler)
	http.Handle("/", mux)

	port := ":8080"
	log.Println("Mipango server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
