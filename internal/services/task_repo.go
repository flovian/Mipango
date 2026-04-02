package services

import (
	"mipango/internal/models"
	"sync"
)

type TaskRepo struct {
	tasks map[string]*models.Task
	mu    sync.Mutex
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		tasks: make(map[string]*models.Task),
	}
}

func (r *TaskRepo) CreateTask(title, objectiveID string, priority int, deadline string) *models.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := generateID() // function below
	task := &models.Task{
		ID:        id,
		Title:     title,
		Objective: objectiveID,
		Priority:  priority,
		Deadline:  deadline,
		Completed: false,
	}
	r.tasks[id] = task
	return task
}

func (r *TaskRepo) GetTasksByObjective(objectiveID string) []*models.Task {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := []*models.Task{}
	for _, t := range r.tasks {
		if t.Objective == objectiveID {
			list = append(list, t)
		}
	}
	return list
}

// Simple ID generator
import "github.com/google/uuid"

func generateID() string {
	return uuid.New().String()
}