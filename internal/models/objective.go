package models

type Objective struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Priority  int    `json:"priority"`
	Completed bool   `json:"completed"`
	Deadline  string `json:"deadline"` // store as string for simplicity
}