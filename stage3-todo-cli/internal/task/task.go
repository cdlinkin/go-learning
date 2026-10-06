package task

import (
	"fmt"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Done     bool   `json:"done"`
	Priority string `json:"priority"`
}

const (
	high   = "high"
	medium = "medium"
	low    = "low"
)

type TaskManager struct {
	Tasks []Task
}

func (t *TaskManager) Add(title, priority string) error {
	if priority != high && priority != medium && priority != low {
		return ErrNoSuchPriority
	}

	var maxID int
	for _, task := range t.Tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	t.Tasks = append(t.Tasks, Task{
		ID:       maxID + 1,
		Title:    title,
		Done:     false,
		Priority: priority,
	})
	return nil
}

func (t *TaskManager) List(showDone bool) {
	countDone := 0
	for _, task := range t.Tasks {
		if showDone == true && task.Done == false {
			continue
		}

		countDone++

		var done string
		if task.Done {
			done = "[ ✔ ]"
		} else {
			done = "[   ]"
		}
		fmt.Printf("%d. %s %s [ %s ]\n", countDone, task.Title, done, task.Priority)
	}
}

func (t *TaskManager) Done(id int) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Done = true
			return nil
		}
	}
	return ErrTaskNotFound
}

func (t *TaskManager) Delete(id int) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return nil
		}
	}
	return ErrTaskNotFound
}
