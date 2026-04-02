package services

import (
	"github.com/google/uuid"
	"mipango/internal/models"
	"mipango/internal/repository"
)

// -- ObjectiveService--
type ObjectiveService struct {
	repo     *repository.ObjectiveRepo
	taskRepo *TaskRepo
}

// ---------------- Constructor ----------------
func NewObjectiveService(repo *repository.ObjectiveRepo, taskRepo *TaskRepo) *ObjectiveService {
	return &ObjectiveService{
		repo:     repo,
		taskRepo: taskRepo,
	}
}

// ---------------- Objective Methods ----------------
func (s *ObjectiveService) CreateObjective(title string, deadline string, priority int) models.Objective {
	id := uuid.New().String()
	obj := models.Objective{
		ID:        id,
		Title:     title,
		Priority:  priority,
		Completed: false,
		Deadline:  deadline, // You can later parse it to time.Time if needed
	}
	s.repo.Save(obj)
	return obj
}

func (s *ObjectiveService) GetAllObjectives() []models.Objective {
	return s.repo.GetAll()
}

// ---------------- Task Methods ----------------
func (s *ObjectiveService) CreateTask(title, objectiveID string, priority int, deadline string) *models.Task {
	return s.taskRepo.CreateTask(title, objectiveID, priority, deadline)
}

func (s *ObjectiveService) GetTasks(objectiveID string) []*models.Task {
	return s.taskRepo.GetTasksByObjective(objectiveID)
}