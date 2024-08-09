package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/ARUMANDESU/todo-app/internal/domain"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"log"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
	"time"
)

//go:embed migrations
var migrationsFs embed.FS

type Storage struct {
	db *sql.DB
}

func getDataSource() string {
	cacheDir, _ := os.UserCacheDir()
	dataDir := filepath.Join(cacheDir, "todo-app")
	os.MkdirAll(dataDir, os.FileMode(0755))

	// if file is not found, it will be created automatically
	if _, err := os.Stat(filepath.Join(dataDir, "tasks.db")); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(dataDir, "tasks.db"))
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	return filepath.Join(dataDir, "tasks.db")
}

func NewStorage() (*Storage, error) {
	if err := migrateSchema(nil); err != nil {
		return nil, fmt.Errorf("failed to perform migrations: %w", err)
	}
	db, err := sql.Open("sqlite", getDataSource())
	if err != nil {
		return nil, fmt.Errorf("open sqlite connection: %w", err)
	}

	return &Storage{db: db}, nil
}

func migrateSchema(nSteps *int) error {
	db, err := sql.Open("sqlite", getDataSource())
	if err != nil {
		return fmt.Errorf("open sqlite connection: %w", err)
	}

	migrateDriver, err := sqlite.WithInstance(db, &sqlite.Config{
		MigrationsTable: "migrations",
	})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}
	srcDriver, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return fmt.Errorf("failed to create migration source driver: %w", err)
	}
	preparedMigrations, err := migrate.NewWithInstance(
		"iofs",
		srcDriver,
		"",
		migrateDriver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration tooling instance: %w", err)
	}
	defer func() {
		preparedMigrations.Close()
		db.Close()
	}()
	if nSteps != nil {
		fmt.Printf("stepping migrations %d...\n", *nSteps)
		err = preparedMigrations.Steps(*nSteps)
	} else {
		err = preparedMigrations.Up()
	}

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("Successfully applied db migrations")
	return nil
}

func (s Storage) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	const op = "storage.sqlite.task.get_all"

	stmt, err := s.db.Prepare(`SELECT id, title, status, priority, due_date, created_at, modified_at, description, tags FROM tasks`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Status,
			&task.Priority,
			&task.DueDate,
			&task.CreatedAt,
			&task.ModifiedAt,
			&task.Description,
			&task.Tags,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s Storage) GetTaskByID(ctx context.Context, id string) (domain.Task, error) {
	const op = "storage.sqlite.task.get_by_id"

	stmt, err := s.db.Prepare(`SELECT id, title, status, priority, due_date, created_at, modified_at, description, tags  FROM tasks WHERE id = ?`)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	var task domain.Task
	err = stmt.QueryRowContext(ctx, id).Scan(
		&task.ID,
		&task.Title,
		&task.Status,
		&task.Priority,
		&task.DueDate,
		&task.CreatedAt,
		&task.ModifiedAt,
		&task.Description,
		&task.Tags,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Task{}, fmt.Errorf("%s: %w", op, domain.ErrTaskNotFound)
		}
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return task, nil
}

func (s Storage) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	const op = "storage.sqlite.task.create"

	stmt, err := s.db.Prepare(`INSERT INTO tasks(id, title, status, priority, due_date, created_at, modified_at) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(
		ctx,
		task.ID,
		task.Title,
		task.Status,
		task.Priority,
		task.DueDate,
		task.CreatedAt,
		task.ModifiedAt,
	)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return task, nil
}

func (s Storage) UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	const op = "storage.sqlite.task.update"

	stmt, err := s.db.Prepare(`UPDATE tasks SET title = ?, status = ?, priority = ?, due_date = ?, modified_at = ?,description = ?, tags =? WHERE id = ?`)
	if err != nil {
		return domain.Task{}, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(
		ctx,
		task.Title,
		task.Status,
		task.Priority,
		task.DueDate,
		time.Now(),
		task.Description,
		task.Tags.Value(),
		task.ID,
	)
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

func (s Storage) DeleteTask(ctx context.Context, id string) error {
	const op = "storage.sqlite.task.delete"

	stmt, err := s.db.Prepare(`DELETE FROM tasks WHERE id = ?`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, id)
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
