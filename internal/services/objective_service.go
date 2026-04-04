package services

import (
	"sort"

	"github.com/google/uuid"
	"mipango/internal/models"
	"mipango/internal/repository"
)

type ObjectiveService struct {
	objRepo  *repository.ObjectiveRepo
	taskRepo *repository.TaskRepo
}

func NewObjectiveService(objRepo *repository.ObjectiveRepo, taskRepo *repository.TaskRepo) *ObjectiveService {
	return &ObjectiveService{
		objRepo:  objRepo,
		taskRepo: taskRepo,
	}
}

// CreateObjective creates a new objective and stores it
func (s *ObjectiveService) CreateObjective(title, deadline string, priority int) (models.Objective, error) {
	id := uuid.New().String()
	obj := models.Objective{
		ID:        id,
		Title:     title,
		Priority:  priority,
		Completed: false,
		Deadline:  deadline,
		Progress:  0,
	}
	err := s.objRepo.Save(obj)
	return obj, err
}

// GetAllObjectives retrieves objectives and calculates their progress
func (s *ObjectiveService) GetAllObjectives() ([]models.Objective, error) {
	objs, err := s.objRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range objs {
		tasks, _ := s.taskRepo.GetTasksByObjective(objs[i].ID)
		if len(tasks) > 0 {
			completed := 0
			for _, t := range tasks {
				if t.Completed {
					completed++
				}
			}
			objs[i].Progress = (completed * 100) / len(tasks)
		} else {
			objs[i].Progress = 0
		}
	}
	return objs, nil
}

// CreateTask saves a new task
func (s *ObjectiveService) CreateTask(title, objectiveID string, priority int, deadline string) (models.Task, error) {
	id := uuid.New().String()
	task := models.Task{
		ID:        id,
		Title:     title,
		Objective: objectiveID,
		Priority:  priority,
		Completed: false,
		Deadline:  deadline,
	}
	err := s.taskRepo.Save(task)
	return task, err
}

// GetTasks fetches all tasks linked to an objective
func (s *ObjectiveService) GetTasks(objectiveID string) ([]models.Task, error) {
	return s.taskRepo.GetTasksByObjective(objectiveID)
}

// UpdateTaskCompletion toggles the task completion status
func (s *ObjectiveService) UpdateTaskCompletion(taskID string, completed bool) error {
	return s.taskRepo.UpdateCompletion(taskID, completed)
}

// GetSmartSuggestion returns the most urgent and important pending task across all objectives
func (s *ObjectiveService) GetSmartSuggestion() (*models.Task, error) {
	objs, err := s.objRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var allPending []models.Task
	for _, o := range objs {
		tasks, _ := s.taskRepo.GetTasksByObjective(o.ID)
		for _, t := range tasks {
			if !t.Completed {
				allPending = append(allPending, t)
			}
		}
	}

	if len(allPending) == 0 {
		return nil, nil // No pending tasks
	}

	// Sort by highest priority, then by earliest deadline
	sort.Slice(allPending, func(i, j int) bool {
		if allPending[i].Priority != allPending[j].Priority {
			return allPending[i].Priority > allPending[j].Priority
		}
		return allPending[i].Deadline < allPending[j].Deadline
	})

	return &allPending[0], nil
}