package internal

type Task struct {
	provider TaskProvider
	modifier TaskModifier
}

type TaskProvider interface {
	GetAll() ([]Task, error)
	GetByID(id int) (Task, error)
}

type TaskModifier interface {
	Create(task Task) (Task, error)
	Update(task Task) (Task, error)
	Delete(id int) error
}

func NewTask(provider TaskProvider, modifier TaskModifier) Task {
	return Task{
		provider: provider,
		modifier: modifier,
	}
}
