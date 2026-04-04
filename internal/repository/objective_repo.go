package repository

import (
	"database/sql"
	"mipango/internal/models"
)

type ObjectiveRepo struct {
	db *sql.DB
}

func NewObjectiveRepo(db *sql.DB) *ObjectiveRepo {
	return &ObjectiveRepo{db: db}
}

func (r *ObjectiveRepo) Save(obj models.Objective) error {
	query := `INSERT INTO objectives (id, title, priority, completed, deadline) VALUES (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, obj.ID, obj.Title, obj.Priority, obj.Completed, obj.Deadline)
	return err
}

func (r *ObjectiveRepo) GetAll() ([]models.Objective, error) {
	query := `SELECT id, title, priority, completed, deadline FROM objectives`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Objective
	for rows.Next() {
		var obj models.Objective
		if err := rows.Scan(&obj.ID, &obj.Title, &obj.Priority, &obj.Completed, &obj.Deadline); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}
	// Return empty slice instead of nil if no items
	if list == nil {
		list = []models.Objective{}
	}
	return list, nil
}