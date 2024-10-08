// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/ARUMANDESU/todo-app/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// TaskProvider is an autogenerated mock type for the TaskProvider type
type TaskProvider struct {
	mock.Mock
}

// GetAllTasks provides a mock function with given fields: ctx
func (_m *TaskProvider) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Task, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: ctx, id
func (_m *TaskProvider) GetTaskByID(ctx context.Context, id string) (domain.Task, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Task); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskProvider creates a new instance of TaskProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskProvider {
	mock := &TaskProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
