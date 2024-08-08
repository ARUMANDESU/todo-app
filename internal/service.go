package internal

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/ARUMANDESU/todo-app/internal/domain"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/gommon/log"
)

type Task struct {
	provider TaskProvider
	modifier TaskModifier
}

//go:generate mockery --name TaskProvider
type TaskProvider interface {
	GetAllTasks() ([]domain.Task, error)
	GetTaskByID(id int) (domain.Task, error)
}

//go:generate mockery --name TaskModifier
type TaskModifier interface {
	CreateTask(task domain.Task) (domain.Task, error)
	UpdateTask(task domain.Task) (domain.Task, error)
	DeleteTask(id int) error
}

func NewTask(provider TaskProvider, modifier TaskModifier) Task {
	return Task{
		provider: provider,
		modifier: modifier,
	}
}

func (t Task) GetAll() ([]domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t Task) GetByID(id int) (domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t Task) Create(request domain.CreateTaskRequest) (domain.Task, error) {
	const op = "service.task.create"
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required, validation.By(validateTitle)),
		validation.Field(&request.Priority, validation.Required),
		validation.Field(&request.DueDate, validation.By(validateDueDate)),
	)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%w: %w", domain.ErrInvalidArguments, err)
	}

	uid, err := uuid.NewUUID()
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	task := domain.Task{
		ID:         uid,
		Title:      strings.Trim(request.Title, " "),
		Status:     domain.TaskStatusTodo,
		Priority:   request.Priority,
		DueDate:    request.DueDate,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	task, err = t.modifier.CreateTask(task)
	if err != nil {
		return domain.Task{}, handleError(op, err)
	}

	return task, nil
}

func (t Task) Update(request domain.UpdateTaskRequest) (domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t Task) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func handleError(op string, err error) error {
	switch {
	case errors.Is(err, domain.ErrTaskNotFound):
		return domain.ErrTaskNotFound
	case errors.Is(err, domain.ErrInvalidArguments):
		return domain.ErrInvalidArguments
	default:
		log.Error(op, err)
		return domain.ErrInternal
	}
}
