package models

import "time"

// Objective represents a user's goal
type Objective struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Deadline  time.Time `json:"deadline"`
	Priority  int       `json:"priority"` // 1 = low, 5 = high
	Completed bool      `json:"completed"`
}