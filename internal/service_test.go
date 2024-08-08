package internal

import (
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

func TestTask_Create(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		suite := newSuite(t)
		defer suite.mockTaskModifier.AssertExpectations(t)

		suite.mockTaskModifier.On("CreateTask", mock.Anything).Return(domain.Task{}, nil)

		_, err := suite.taskService.Create(domain.CreateTaskRequest{
			Title:    "test",
			Status:   domain.TaskStatusTodo,
			Priority: domain.TaskPriorityLow,
			DueDate:  time.Now().AddDate(0, 0, 1),
		})

		assert.NoError(t, err)
	})
}

func TestTask_Create_FailPath(t *testing.T) {
	t.Run("fail validation", func(t *testing.T) {
		suite := newSuite(t)
		defer suite.mockTaskModifier.AssertExpectations(t)

		_, err := suite.taskService.Create(domain.CreateTaskRequest{})

		assert.Error(t, err)
	})

	t.Run("fail create task", func(t *testing.T) {
		suite := newSuite(t)
		defer suite.mockTaskModifier.AssertExpectations(t)

		suite.mockTaskModifier.On("CreateTask", mock.Anything).Return(domain.Task{}, domain.ErrInternal)

		_, err := suite.taskService.Create(domain.CreateTaskRequest{
			Title:    "test",
			Status:   domain.TaskStatusTodo,
			Priority: domain.TaskPriorityLow,
			DueDate:  time.Now().AddDate(0, 0, 1),
		})

		assert.Error(t, err)
	})
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
