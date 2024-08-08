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

func TestTask_Create_Valid(t *testing.T) {
	suite := newSuite(t)
	ctx := context.Background()
	dueDate := time.Now().Add(24 * time.Hour)
	request := domain.CreateTaskRequest{
		Title:    "New Task",
		Priority: domain.TaskPriorityHigh,
		DueDate:  &dueDate,
	}

	suite.mockTaskModifier.On("CreateTask", ctx, mock.AnythingOfType("domain.Task")).Return(domain.Task{}, nil)

	_, err := suite.taskService.Create(ctx, request)

	assert.NoError(t, err)
}

func TestTask_Create_InvalidRequest(t *testing.T) {
	suite := newSuite(t)

	suite.mockTaskModifier.On("CreateTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(domain.Task{}, assert.AnError)

	ctx := context.Background()
	dueDate := time.Now().Add(24 * time.Hour)
	request := domain.CreateTaskRequest{
		Title:    "test-title",
		Priority: "High",
		DueDate:  &dueDate,
	}

	_, err := suite.taskService.Create(ctx, request)

	assert.Error(t, err)

	suite.mockTaskModifier.AssertCalled(t, "CreateTask", mock.Anything, mock.AnythingOfType("domain.Task"))
}

func TestTask_Update_Valid(t *testing.T) {
	dueDateOld := time.Now().Add(24 * time.Hour)
	dueDateNew := time.Now().Add(48 * time.Hour)
	tests := []struct {
		name     string
		request  domain.UpdateTaskRequest
		setup    func(suite *Suite)
		expected domain.Task
	}{
		{
			name: "update title and priority",
			request: domain.UpdateTaskRequest{
				ID:       "123",
				Title:    "Updated Task",
				Priority: domain.TaskPriorityMedium,
				DueDate:  &dueDateNew,
			},
			setup: func(suite *Suite) {
				suite.mockTaskProvider.On("GetTaskByID", mock.Anything, "123").Return(domain.Task{
					ID:         "123",
					Title:      "Old Task",
					Status:     domain.TaskStatusTodo,
					Priority:   domain.TaskPriorityHigh,
					DueDate:    &dueDateOld,
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}, nil)
				suite.mockTaskModifier.On("UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(domain.Task{
					ID:         "123",
					Title:      "Updated Task",
					Status:     domain.TaskStatusTodo,
					Priority:   domain.TaskPriorityMedium,
					DueDate:    &dueDateNew,
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}, nil)
			},
			expected: domain.Task{
				ID:         "123",
				Title:      "Updated Task",
				Status:     domain.TaskStatusTodo,
				Priority:   domain.TaskPriorityMedium,
				DueDate:    &dueDateNew,
				CreatedAt:  time.Now(),
				ModifiedAt: time.Now(),
			},
		},
		{
			name: "update status",
			request: domain.UpdateTaskRequest{
				ID:     "123",
				Status: domain.TaskStatusDone,
			},
			setup: func(suite *Suite) {
				suite.mockTaskProvider.On("GetTaskByID", mock.Anything, "123").Return(domain.Task{
					ID:         "123",
					Title:      "Old Task",
					Status:     domain.TaskStatusTodo,
					Priority:   domain.TaskPriorityHigh,
					DueDate:    &dueDateOld,
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}, nil)
				suite.mockTaskModifier.On("UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(domain.Task{
					ID:         "123",
					Title:      "Old Task",
					Status:     domain.TaskStatusDone,
					Priority:   domain.TaskPriorityHigh,
					DueDate:    &dueDateOld,
					CreatedAt:  time.Now(),
					ModifiedAt: time.Now(),
				}, nil)
			},
			expected: domain.Task{
				ID:         "123",
				Title:      "Old Task",
				Status:     domain.TaskStatusDone,
				Priority:   domain.TaskPriorityHigh,
				DueDate:    &dueDateOld,
				CreatedAt:  time.Now(),
				ModifiedAt: time.Now(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suite := newSuite(t)
			ctx := context.Background()
			tt.setup(suite)

			task, err := suite.taskService.Update(ctx, tt.request)

			assert.NoError(t, err)
			isTaskEqual(t, tt.expected, task, time.Millisecond)
		})
	}
}

func TestTask_Update_Invalid(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		suite := newSuite(t)

		suite.mockTaskProvider.On("GetTaskByID", mock.Anything, mock.AnythingOfType("string")).Return(domain.Task{}, domain.ErrTaskNotFound)

		ctx := context.Background()
		dueDate := time.Now().Add(24 * time.Hour)
		request := domain.UpdateTaskRequest{
			ID:       "test-id",
			Title:    "Updated Task",
			Priority: "Medium",
			DueDate:  &dueDate,
		}

		_, err := suite.taskService.Update(ctx, request)

		assert.Error(t, err)

		suite.mockTaskProvider.AssertCalled(t, "GetTaskByID", mock.Anything, mock.AnythingOfType("string"))
		suite.mockTaskModifier.AssertNotCalled(t, "UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task"))
	})

	t.Run("invalid request", func(t *testing.T) {
		suite := newSuite(t)

		suite.mockTaskProvider.On("GetTaskByID", mock.Anything, mock.AnythingOfType("string")).Return(domain.Task{}, nil)
		suite.mockTaskModifier.On("UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(domain.Task{}, assert.AnError)

		ctx := context.Background()
		dueDate := time.Now().Add(24 * time.Hour)
		request := domain.UpdateTaskRequest{
			ID:       "123",
			Title:    "",
			Priority: "Medium",
			DueDate:  &dueDate,
		}

		_, err := suite.taskService.Update(ctx, request)

		assert.Error(t, err)

		suite.mockTaskModifier.AssertCalled(t, "UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task"))
	})
}

func TestTask_Update_InvalidArguments(t *testing.T) {
	t.Run("title less than 3 character", func(t *testing.T) {
		suite := newSuite(t)

		ctx := context.Background()
		dueDate := time.Now().Add(24 * time.Hour)
		request := domain.UpdateTaskRequest{
			ID:       "123",
			Title:    "12",
			Priority: "Medium",
			DueDate:  &dueDate,
		}

		_, err := suite.taskService.Update(ctx, request)

		assert.Error(t, err)

		suite.mockTaskProvider.AssertNotCalled(t, "GetTaskByID", mock.Anything, mock.AnythingOfType("string"))
		suite.mockTaskModifier.AssertNotCalled(t, "UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task"))
	})

	t.Run("invalid due date", func(t *testing.T) {
		suite := newSuite(t)

		ctx := context.Background()
		dueDate := time.Now().Add(-24 * time.Hour)
		request := domain.UpdateTaskRequest{
			ID:       "123",
			Title:    "Updated Task",
			Priority: "Medium",
			DueDate:  &dueDate,
		}

		_, err := suite.taskService.Update(ctx, request)

		assert.Error(t, err)

		suite.mockTaskProvider.AssertNotCalled(t, "GetTaskByID", mock.Anything, mock.AnythingOfType("string"))
		suite.mockTaskModifier.AssertNotCalled(t, "UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task"))
	})

	t.Run("empty id", func(t *testing.T) {
		suite := newSuite(t)

		ctx := context.Background()
		dueDate := time.Now().Add(24 * time.Hour)
		request := domain.UpdateTaskRequest{
			ID:       "",
			Title:    "Updated Task",
			Priority: "Medium",
			DueDate:  &dueDate,
		}

		_, err := suite.taskService.Update(ctx, request)

		assert.Error(t, err)

		suite.mockTaskProvider.AssertNotCalled(t, "GetTaskByID", mock.Anything, mock.AnythingOfType("string"))
		suite.mockTaskModifier.AssertNotCalled(t, "UpdateTask", mock.Anything, mock.AnythingOfType("domain.Task"))
	})
}

func isTaskEqual(t *testing.T, expected, actual domain.Task, tolerance time.Duration) {
	t.Helper()
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Status, actual.Status)
	assert.Equal(t, expected.Priority, actual.Priority)
	assert.WithinDuration(t, *expected.DueDate, *actual.DueDate, tolerance)
	assert.WithinDuration(t, expected.CreatedAt, actual.CreatedAt, tolerance)
	assert.WithinDuration(t, expected.ModifiedAt, actual.ModifiedAt, tolerance)

}
