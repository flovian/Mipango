package services

import (
	"mipango/internal/models"
	"sync"

	"github.com/google/uuid"
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

	task := &models.Task{
		ID:        uuid.New().String(),
		Title:     title,
		Objective: objectiveID,
		Priority:  priority,
		Deadline:  deadline,
		Completed: false,
	}
	r.tasks[task.ID] = task
	return task
}