package task

import (
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskManager struct {
	Tasks []Task
}

func (t *TaskManager) Add(title string) {
	t.Tasks = append(t.Tasks, Task{
		ID:    len(t.Tasks) + 1,
		Title: title,
		Done:  false,
	})
}

func (t *TaskManager) List() {
	for _, task := range t.Tasks {
		var done string
		if task.Done {
			done = "[x]"
		} else {
			done = "[ ]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, task.Title, done)
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
