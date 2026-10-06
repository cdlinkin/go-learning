package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"stage3-todo-cli/internal/task"
)

type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{
		path: path,
	}
}

func (f *FileStorage) Save(tasks []task.Task) error {
	tasksBytes, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}

	if err := os.WriteFile(f.path, tasksBytes, 0644); err != nil {
		return fmt.Errorf("write tasks file: %w", err)
	}
	return nil
}

func (f *FileStorage) Load() ([]task.Task, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var tasks []task.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("unmarshal tasks: %w", err)
	}

	return tasks, nil
}
