package internal

import "github.com/ARUMANDESU/todo-app/internal/domain"

type Task struct {
	provider TaskProvider
	modifier TaskModifier
}

type TaskProvider interface {
	GetAll() ([]domain.Task, error)
	GetByID(id int) (domain.Task, error)
}

type TaskModifier interface {
	Create(task domain.Task) (domain.Task, error)
	Update(task domain.Task) (domain.Task, error)
	Delete(id int) error
}

func NewTask(provider TaskProvider, modifier TaskModifier) Task {
	return Task{
		provider: provider,
		modifier: modifier,
	}
}
