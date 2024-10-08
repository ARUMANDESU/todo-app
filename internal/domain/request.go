package domain

import "time"

type CreateTaskRequest struct {
	Title    string       `json:"title"`
	Priority TaskPriority `json:"priority"`
	DueDate  *time.Time   `json:"due_date"`
}

type UpdateTaskRequest struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description *string      `json:"description"` // pointer to make it optional
	Tags        []string     `json:"tags"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	DueDate     *time.Time   `json:"due_date"` // pointer to make it optional
}
