package models

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Objective string `json:"objective"`
	Priority  int    `json:"priority"`
	Completed bool   `json:"completed"`
	Deadline  string `json:"deadline"`
}