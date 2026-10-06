package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"stage3-todo-cli/internal/storage"
	"stage3-todo-cli/internal/task"
	"strconv"
)

func main() {
	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("user home dir: %v", err)
		return
	}
	path = filepath.Join(path, "tasks.json")

	fileStorage := storage.NewFileStorage(path)

	tasks, err := fileStorage.Load()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			tasks = []task.Task{}
		} else {
			fmt.Printf("file storage load: %v", err)
			return
		}
	}

	taskManager := task.TaskManager{
		Tasks: tasks,
	}

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			return
		}

		taskManager.Add(os.Args[2])
	case "list":
		taskManager.List()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("strconv id: %v", err)
			return
		}

		if err := taskManager.Done(id); err != nil {
			fmt.Printf("task manager done: %v", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("strconv id: %v", err)
			return
		}

		if err := taskManager.Delete(id); err != nil {
			fmt.Printf("task manager delete: %v", err)
		}
	default:
		fmt.Printf("your command is unknown: %s\n", os.Args[1])
	}

	if err := fileStorage.Save(taskManager.Tasks); err != nil {
		fmt.Printf("file storage save: %v", err)
		return
	}
}
