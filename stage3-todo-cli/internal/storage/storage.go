package storage

import "stage3-todo-cli/internal/task"

type Storage interface {
	Load() ([]task.Task, error)
	Save([]task.Task) error
}
