package domain

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID         uuid.UUID    `json:"id"`
	Title      string       `json:"title"`
	Status     TaskStatus   `json:"status"`
	Priority   TaskPriority `json:"priority"`
	DueDate    *time.Time   `json:"due_date,omitempty"`
	CreatedAt  time.Time    `json:"created_at"`
	ModifiedAt time.Time    `json:"modified_at"`
}

type TaskStatus string
type TaskPriority string

const (
	TaskStatusTodo TaskStatus = "todo"
	TaskStatusDone TaskStatus = "done"
)

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
)

var AllTaskStatus = []struct {
	Value  TaskStatus
	TSName string
}{
	{TaskStatusTodo, "TODO"},
	{TaskStatusDone, "DONE"},
}

var AllTaskPriority = []struct {
	Value  TaskPriority
	TPName string
}{
	{TaskPriorityLow, "LOW"},
	{TaskPriorityMedium, "MEDIUM"},
	{TaskPriorityHigh, "HIGH"},
}
