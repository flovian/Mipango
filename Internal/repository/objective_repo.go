package repository

import (
	"errors"
	"sync"

	"mipango/internal/models"
)

var (
	ErrObjectiveNotFound = errors.New("objective not found")
)

// In-memory store
type ObjectiveRepo struct {
	data map[string]models.Objective
	mu   sync.RWMutex
}

func NewObjectiveRepo() *ObjectiveRepo {
	return &ObjectiveRepo{
		data: make(map[string]models.Objective),
	}
}

func (r *ObjectiveRepo) Save(obj models.Objective) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[obj.ID] = obj
}

func (r *ObjectiveRepo) GetAll() []models.Objective {
	r.mu.RLock()
	defer r.mu.RUnlock()
	objs := make([]models.Objective, 0, len(r.data))
	for _, obj := range r.data {
		objs = append(objs, obj)
	}
	return objs
}

func (r *ObjectiveRepo) GetByID(id string) (models.Objective, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	obj, ok := r.data[id]
	if !ok {
		return models.Objective{}, ErrObjectiveNotFound
	}
	return obj, nil
}