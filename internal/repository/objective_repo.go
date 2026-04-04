package repository

import (
	"mipango/internal/models"
	"sync"
)

type ObjectiveRepo struct {
	data map[string]models.Objective
	mu   sync.Mutex
}

func NewObjectiveRepo() *ObjectiveRepo {
	return &ObjectiveRepo{data: make(map[string]models.Objective)}
}

func (r *ObjectiveRepo) Save(obj models.Objective) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[obj.ID] = obj
}

func (r *ObjectiveRepo) GetAll() []models.Objective {
	r.mu.Lock()
	defer r.mu.Unlock()
	list := make([]models.Objective, 0, len(r.data))
	for _, obj := range r.data {
		list = append(list, obj)
	}
	return list
}