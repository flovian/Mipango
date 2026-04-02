package models

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Objective string `json:"objective_id"`
	Priority  int    `json:"priority"`
	Deadline  string `json:"deadline"`
	Completed bool   `json:"completed"`
}