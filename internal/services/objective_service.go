package services

import (
	"github.com/google/uuid"
	"mipango/internal/models"
	"mipango/internal/repository"
)

type ObjectiveService struct {
	objRepo  *repository.ObjectiveRepo
	taskRepo *TaskRepo
}

func NewObjectiveService(objRepo *repository.ObjectiveRepo, taskRepo *TaskRepo) *ObjectiveService {
	return &ObjectiveService{
		objRepo:  objRepo,
		taskRepo: taskRepo,
	}
}

// Objective methods
func (s *ObjectiveService) CreateObjective(title, deadline string, priority int) models.Objective {
	id := uuid.New().String()
	obj := models.Objective{
		ID:        id,
		Title:     title,
		Priority:  priority,
		Deadline:  deadline,
		Completed: false,
	}
	s.objRepo.Save(obj)
	return obj
}

func (s *ObjectiveService) GetAllObjectives() []models.Objective {
	return s.objRepo.GetAll()
}

// Task methods
func (s *ObjectiveService) CreateTask(title, objectiveID string, priority int, deadline string) *models.Task {
	return s.taskRepo.CreateTask(title, objectiveID, priority, deadline)
}

func (s *ObjectiveService) GetTasks(objectiveID string) []*models.Task {
	return s.taskRepo.GetTasksByObjective(objectiveID)
}