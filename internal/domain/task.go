package domain

import "time"

type Task struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Status    TaskStatus   `json:"status"`
	Priority  TaskPriority `json:"priority"`
	Timestamp time.Time    `json:"timestamp"`
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
