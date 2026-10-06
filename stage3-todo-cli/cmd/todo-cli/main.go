package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"stage3-todo-cli/internal/storage"
	"stage3-todo-cli/internal/task"
	"strconv"
)

func main() {
	filePath := flag.String("file", "", "path to tasks file")
	flag.Parse()

	var path string

	if *filePath != "" {
		path = *filePath
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("user home dir: %v\n", err)
			return
		}

		path = filepath.Join(homeDir, "tasks.json")
	}

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

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("not enough arguments")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("not enough arguments")
			return
		}

		taskManager.Add(args[1])
	case "list":
		taskManager.List()
	case "done":
		if len(args) < 2 {
			fmt.Println("not enough arguments")
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("strconv id: %v", err)
			return
		}

		if err := taskManager.Done(id); err != nil {
			if errors.Is(err, task.ErrTaskNotFound) {
				fmt.Println(task.ErrTaskNotFound)
			} else {
				fmt.Println(err)
			}
		}
	case "delete":
		if len(args) < 2 {
			fmt.Println("not enough arguments")
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("strconv id: %v", err)
			return
		}

		if err := taskManager.Delete(id); err != nil {
			if errors.Is(err, task.ErrTaskNotFound) {
				fmt.Println(task.ErrTaskNotFound)
			} else {
				fmt.Println(err)
			}
		}
	default:
		fmt.Printf("your command is unknown: %s\n", args[0])
	}

	if err := fileStorage.Save(taskManager.Tasks); err != nil {
		fmt.Printf("file storage save: %v", err)
		return
	}
}
