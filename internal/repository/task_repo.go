package repository

import (
	"database/sql"
	"mipango/internal/models"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) Save(task models.Task) error {
	query := `INSERT INTO tasks (id, title, objective_id, priority, completed, deadline) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, task.ID, task.Title, task.Objective, task.Priority, task.Completed, task.Deadline)
	return err
}

func (r *TaskRepo) UpdateCompletion(id string, completed bool) error {
	query := `UPDATE tasks SET completed = ? WHERE id = ?`
	_, err := r.db.Exec(query, completed, id)
	return err
}

func (r *TaskRepo) GetTasksByObjective(objectiveID string) ([]models.Task, error) {
	query := `SELECT id, title, objective_id, priority, completed, deadline FROM tasks WHERE objective_id = ?`
	rows, err := r.db.Query(query, objectiveID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Objective, &t.Priority, &t.Completed, &t.Deadline); err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	if list == nil {
		list = []models.Task{}
	}
	return list, nil
}
