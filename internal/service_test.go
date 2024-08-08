package internal

import (
	"context"
	"github.com/ARUMANDESU/todo-app/internal/domain"
	"github.com/ARUMANDESU/todo-app/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type Suite struct {
	mockTaskProvider *mocks.TaskProvider
	mockTaskModifier *mocks.TaskModifier
	taskService      Task
}

func newSuite(t *testing.T) *Suite {
	mockTaskProvider := mocks.NewTaskProvider(t)
	mockTaskModifier := mocks.NewTaskModifier(t)
	taskService := NewTask(mockTaskProvider, mockTaskModifier)
	return &Suite{
		mockTaskProvider: mockTaskProvider,
		mockTaskModifier: mockTaskModifier,
		taskService:      taskService,
	}
}

func TestCreate_ValidRequest_CreatesTask(t *testing.T) {
	suite := newSuite(t)
	ctx := context.Background()
	dueDate := time.Now().Add(24 * time.Hour)
	request := domain.CreateTaskRequest{
		Title:    "New Task",
		Priority: "High",
		DueDate:  &dueDate,
	}

	suite.mockTaskModifier.On("CreateTask", ctx, mock.AnythingOfType("domain.Task")).Return(domain.Task{
		ID:         "123",
		Title:      "New Task",
		Status:     domain.TaskStatusTodo,
		Priority:   "High",
		DueDate:    request.DueDate,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}, nil)

	task, err := suite.taskService.Create(ctx, request)

	assert.NoError(t, err)
	assert.Equal(t, "New Task", task.Title)
	assert.Equal(t, "High", task.Priority)
	assert.Equal(t, domain.TaskStatusTodo, task.Status)
}

func TestCreate_InvalidRequest_ReturnsError(t *testing.T) {
	suite := newSuite(t)
	ctx := context.Background()
	dueDate := time.Now().Add(24 * time.Hour)
	request := domain.CreateTaskRequest{
		Title:    "",
		Priority: "High",
		DueDate:  &dueDate,
	}

	task, err := suite.taskService.Create(ctx, request)

	assert.Error(t, err)
	assert.Equal(t, domain.Task{}, task)
}

func TestUpdate_ValidRequest_UpdatesTask(t *testing.T) {
	suite := newSuite(t)
	ctx := context.Background()
	dueDate := time.Now().Add(24 * time.Hour)
	request := domain.UpdateTaskRequest{
		ID:       "123",
		Title:    "Updated Task",
		Priority: "Medium",
		DueDate:  &dueDate,
	}

	suite.mockTaskProvider.On("GetTaskByID", ctx, "123").Return(domain.Task{
		ID:         "123",
		Title:      "Old Task",
		Status:     domain.TaskStatusTodo,
		Priority:   "High",
		DueDate:    &dueDate,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}, nil)

	suite.mockTaskModifier.On("UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(domain.Task{
		ID:         "123",
		Title:      "Updated Task",
		Status:     domain.TaskStatusTodo,
		Priority:   "Medium",
		DueDate:    request.DueDate,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}, nil)

	task, err := suite.taskService.Update(ctx, request)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", task.Title)
	assert.Equal(t, "Medium", task.Priority)
}

func TestUpdate_InvalidRequest_ReturnsError(t *testing.T) {
	suite := newSuite(t)
	ctx := context.Background()
	dueDate := time.Now().Add(24 * time.Hour)
	request := domain.UpdateTaskRequest{
		ID:       "",
		Title:    "Updated Task",
		Priority: "Medium",
		DueDate:  &dueDate,
	}

	task, err := suite.taskService.Update(ctx, request)

	assert.Error(t, err)
	assert.Equal(t, domain.Task{}, task)
}

func isTaskEqual(t *testing.T, expected, actual domain.Task) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.Priority, actual.Priority)
	assert.Equal(t, expected.DueDate, actual.DueDate)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
	assert.Equal(t, expected.ModifiedAt, actual.ModifiedAt)
}
