package internal

import (
	"context"
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
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (domain.Task, error)
}

//go:generate mockery --name TaskModifier
type TaskModifier interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

func NewTask(provider TaskProvider, modifier TaskModifier) Task {
	return Task{
		provider: provider,
		modifier: modifier,
	}
}

func (t Task) GetAll(ctx context.Context) ([]domain.Task, error) {
	const op = "service.task.get_all"

	tasks, err := t.provider.GetAllTasks(ctx)
	if err != nil {
		return nil, handleError(op, err)
	}

	return tasks, nil
}

func (t Task) GetByID(ctx context.Context, id string) (domain.Task, error) {
	const op = "service.task.get_by_id"

	task, err := t.provider.GetTaskByID(ctx, id)
	if err != nil {
		return domain.Task{}, handleError(op, err)
	}

	return task, nil
}

func (t Task) Create(ctx context.Context, request domain.CreateTaskRequest) (domain.Task, error) {
	const op = "service.task.create"
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required, validation.By(validateTitle)),
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
		ID:         uid.String(),
		Title:      strings.Trim(request.Title, " "),
		Status:     domain.TaskStatusTodo,
		Priority:   request.Priority,
		DueDate:    request.DueDate,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	task, err = t.modifier.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, handleError(op, err)
	}

	return task, nil
}

func (t Task) Update(ctx context.Context, request domain.UpdateTaskRequest) (domain.Task, error) {
	const op = "service.task.update"
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ID, validation.Required),
		validation.Field(&request.Title, validation.By(validateTitle)),
		validation.Field(&request.DueDate, validation.By(validateDueDate)),
	)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%w: %w", domain.ErrInvalidArguments, err)
	}

	task, err := t.provider.GetTaskByID(ctx, request.ID)
	if err != nil {
		return domain.Task{}, handleError(op, err)
	}

	if request.Title != "" {
		task.Title = strings.Trim(request.Title, " ")
	}
	if request.Status != "" {
		task.Status = request.Status
	}
	if request.Priority != "" {
		task.Priority = request.Priority
	}
	if request.DueDate != nil {
		task.DueDate = request.DueDate
	}

	task.ModifiedAt = time.Now()

	updateCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	task, err = t.modifier.UpdateTask(updateCtx, task)
	if err != nil {
		return domain.Task{}, handleError(op, err)
	}

	return task, nil
}

func (t Task) Delete(ctx context.Context, id string) error {
	const op = "service.task.delete"

	deleteCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := t.modifier.DeleteTask(deleteCtx, id)
	if err != nil {
		return handleError(op, err)
	}

	return nil
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
