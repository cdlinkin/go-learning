package task

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	manager := TaskManager{}

	err := manager.Add("hello", "high", time.Time{})
	if err != nil {
		t.Fatalf("task not added: %v", err)
	}

	if len(manager.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(manager.Tasks))
	}

	task := manager.Tasks[0]

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}

	if task.Title != "hello" {
		t.Errorf("expected title 'hello', got %q", task.Title)
	}

	if task.Priority != "high" {
		t.Errorf("expected priority 'high', got %q", task.Priority)
	}

	if task.Done {
		t.Error("new task should not be done")
	}

	if !task.Deadline.IsZero() {
		t.Error("new task should not have a deadline")
	}
}

func TestDone(t *testing.T) {
	manager := TaskManager{}

	if err := manager.Add("learn Go", "high", time.Time{}); err != nil {
		t.Fatalf("add task: %v", err)
	}

	err := manager.Done(1)
	if err != nil {
		t.Fatalf("done task: %v", err)
	}

	if !manager.Tasks[0].Done {
		t.Error("task should be done")
	}
}

func TestDelete(t *testing.T) {
	manager := TaskManager{}

	if err := manager.Add("first", "high", time.Time{}); err != nil {
		t.Fatalf("add first task: %v", err)
	}

	if err := manager.Add("second", "medium", time.Time{}); err != nil {
		t.Fatalf("add second task: %v", err)
	}

	err := manager.Delete(1)
	if err != nil {
		t.Fatalf("delete task: %v", err)
	}

	if len(manager.Tasks) != 1 {
		t.Fatalf("expected 1 task after delete, got %d", len(manager.Tasks))
	}

	if manager.Tasks[0].Title != "second" {
		t.Errorf(
			"expected remaining task 'second', got %q",
			manager.Tasks[0].Title,
		)
	}
}
