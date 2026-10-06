package task

import "errors"

var (
	ErrTaskNotFound   = errors.New("Task not found")
	ErrNoSuchPriority = errors.New("There is no such priority")
)
