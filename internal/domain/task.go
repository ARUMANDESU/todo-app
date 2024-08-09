package domain

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		return nil // case when value from the db was NULL
	}
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to cast value to string: %v", value)
	}
	tags := strings.Split(s, ",")
	if len(tags) == 0 {
		return nil
	}
	*a = tags
	return nil
}

func (a StringArray) Value() driver.Value {
	return strings.Join(a, ",")
}

type Task struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Tags        StringArray  `json:"tags"`
	Status      TaskStatus   `json:"status"`
	Priority    TaskPriority `json:"priority"`
	DueDate     *time.Time   `json:"due_date,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	ModifiedAt  time.Time    `json:"modified_at"`
}

type TaskStatus string
type TaskPriority string

const (
	TaskStatusTodo TaskStatus = "todo"
	TaskStatusDone TaskStatus = "done"
)

const (
	TaskPriorityNone   TaskPriority = "none"
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
	TSName string
}{
	{TaskPriorityNone, "NONE"},
	{TaskPriorityLow, "LOW"},
	{TaskPriorityMedium, "MEDIUM"},
	{TaskPriorityHigh, "HIGH"},
}
