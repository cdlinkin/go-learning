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
	"time"
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
		printUsage()
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 3 {
			printUsage()
			return
		}

		var deadline time.Time

		if len(args) == 4 {
			var err error

			deadline, err = time.Parse("2006-01-02", args[3])
			if err != nil {
				fmt.Println("invalid deadline: use YYYY-MM-DD")
				return
			}
		}

		if err := taskManager.Add(args[1], args[2], deadline); err != nil {
			if errors.Is(err, task.ErrNoSuchPriority) {
				fmt.Println("invalid priority: use high, medium or low")
				return
			}
			fmt.Printf("task manager add: %v", err)
			return
		}
	case "list":
		if len(args) > 1 {
			if args[1] == "--done" {
				taskManager.List(true)
			}
		}

		taskManager.List(false)
	case "done":
		if len(args) < 2 {
			printUsage()
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
			printUsage()
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

func printUsage() {
	fmt.Println(`
	- Usage:
  	 	todo add "task title" <priority> [deadline]

	- Priority:
  	 	high
  	 	medium
  	 	low

	- Example:
  	 	todo add "learn Go" high 2026-10-10
`)
}
