package main

import (
	"context"
	"fmt"
	"github.com/ARUMANDESU/todo-app/internal"
	"github.com/ARUMANDESU/todo-app/internal/domain"
	"github.com/ARUMANDESU/todo-app/internal/storage/sqlite"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	taskService TaskService
}

type TaskService interface {
	GetAll() ([]domain.Task, error)
	GetByID(id int) (domain.Task, error)
	Create(request domain.CreateTaskRequest) (domain.Task, error)
	Update(request domain.UpdateTaskRequest) (domain.Task, error)
	Delete(id int) error
}

// NewApp creates a new App application struct
func NewApp() *App {
	sqliteDB := sqlite.NewStorage()
	taskService := internal.NewTask(sqliteDB, sqliteDB)

	return &App{taskService: taskService}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetAllTasks() ([]domain.Task, error) {
	return a.taskService.GetAll()
}

func (a *App) GetTaskByID(id int) (domain.Task, error) {
	return a.taskService.GetByID(id)
}

func (a *App) CreateTask(request domain.CreateTaskRequest) (domain.Task, error) {
	return a.taskService.Create(request)
}

func (a *App) UpdateTask(request domain.UpdateTaskRequest) (domain.Task, error) {
	return a.taskService.Update(request)
}

func (a *App) DeleteTask(id int) error {
	if !a.confirmDeleteTask() {
		return nil
	}

	return a.taskService.Delete(id)
}

func (a *App) confirmDeleteTask() bool {
	response, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "Confirm Deletion",
		Message:       "Are you sure you want to delete this task?",
		Buttons:       []string{"Yes", "No"},
		DefaultButton: "No",
	})
	if err != nil {
		fmt.Println("Error showing dialog:", err)
		return false
	}
	return response == "Yes"
}
