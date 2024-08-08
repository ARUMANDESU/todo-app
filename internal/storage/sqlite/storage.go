package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ARUMANDESU/todo-app/internal/domain"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	dbPath := os.Getenv("SQLITE_DB_PATH")
	if dbPath == "" {
		dbPath = "./db/sqlite/tasks.db" // default path
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil
	}
	return &Storage{db: db}
}

func (s Storage) GetAll() ([]domain.Task, error) {
	const op = "storage.sqlite.task.get_all"

	stmt, err := s.db.Prepare(`SELECT id, title, status, priority, due_date, created_at, modified_at FROM tasks`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.Priority, &task.DueDate, &task.CreatedAt, &task.ModifiedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s Storage) GetByID(id int) (domain.Task, error) {
	const op = "storage.sqlite.task.get_by_id"

	stmt, err := s.db.Prepare(`SELECT id, title, status, priority, due_date, created_at, modified_at  FROM tasks WHERE id = ?`)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	var task domain.Task
	err = stmt.QueryRow(id).Scan(&task.ID, &task.Title, &task.Status, &task.Priority, &task.DueDate, &task.CreatedAt, &task.ModifiedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("%s: %w", op, domain.ErrTaskNotFound)
		}
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return task, nil
}

func (s Storage) Create(task domain.Task) (domain.Task, error) {
	const op = "storage.sqlite.task.create"

	stmt, err := s.db.Prepare(`INSERT INTO tasks(id, title, status, priority, due_date, created_at, modified_at) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(task.ID, task.Title, task.Status, task.Priority, task.DueDate, task.CreatedAt, task.ModifiedAt)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return task, nil
}

func (s Storage) Update(task domain.Task) (domain.Task, error) {
	const op = "storage.sqlite.task.update"

	stmt, err := s.db.Prepare(`UPDATE tasks SET title = ?, status = ?, priority = ?, due_date = ?, modified_at = ? WHERE id = ?`)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(task.Title, task.Status, task.Priority, task.DueDate, task.ModifiedAt, task.ID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return domain.Task{}, fmt.Errorf("%s: %w", op, domain.ErrTaskNotFound)
	}

	return task, nil
}

func (s Storage) Delete(id int) error {
	const op = "storage.sqlite.task.delete"

	stmt, err := s.db.Prepare(`DELETE FROM tasks WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, domain.ErrTaskNotFound)
	}

	return nil
}
