package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Done     bool      `json:"done"`
	Priority string    `json:"priority"`
	Deadline time.Time `json:"deadline"`
}

const (
	high   = "high"
	medium = "medium"
	low    = "low"
)

type TaskManager struct {
	Tasks []Task
}

func (t *TaskManager) Add(title, priority string, deadline time.Time) error {
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
		Deadline: deadline,
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

		deadline := ""

		if !task.Deadline.IsZero() {
			deadline = task.Deadline.Format("2006-01-02")
		}

		var done string
		if task.Done {
			done = "[ ✔ ]"
		} else {
			done = "[   ]"
		}
		fmt.Printf("%d. %s %s [ %s ] | DEADLINE: %s\n", countDone, task.Title, done, task.Priority, deadline)
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
