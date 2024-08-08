package sqlite

import (
	"database/sql"
	"github.com/ARUMANDESU/todo-app/internal"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s Storage) GetAll() ([]internal.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetByID(id int) (internal.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) Create(task internal.Task) (internal.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) Update(task internal.Task) (internal.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
