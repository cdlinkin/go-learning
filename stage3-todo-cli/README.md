# Todo CLI

A small command-line todo app written in Go.

I built this project to practice Go project structure, JSON, file storage, CLI arguments, error handling, interfaces, and unit tests.

## Features

* Add tasks with priority and optional deadline
* List all tasks
* List completed tasks with `--done`
* Mark tasks as completed
* Delete tasks
* Store tasks in a JSON file
* Use a custom file with the `-file` flag

## Usage

Add a task:

```bash
todo add "learn Go" high 2026-10-10
```

Add a task without a deadline:

```bash
todo add "buy some bread" low
```

List tasks:

```bash
todo list
```

Show completed tasks:

```bash
todo list --done
```

Mark a task as done:

```bash
todo done 1
```

Delete a task:

```bash
todo delete 2
```

Available priorities:

```text
high
medium
low
```

## Storage

By default, tasks are stored in `tasks.json` in the user's home directory.

You can also specify a custom file:

```bash
todo -file ./tasks.json list
```

## Project Structure

```text
stage3-todo-cli/
├── cmd/
│   └── todo/
│       └── main.go
├── internal/
│   ├── task/
│   │   ├── task.go
│   │   ├── errors.go
│   │   └── task_test.go
│   └── storage/
│       ├── storage.go
│       └── file.go
├── README.md
└── go.mod
```

* `cmd/todo` — CLI entry point
* `internal/task` — task model and logic
* `internal/storage` — file storage
* `task_test.go` — unit tests

## Tests

The project has unit tests for the main task operations, including adding, completing, deleting tasks, priorities, deadlines, errors, and unique IDs.

Run tests:

```bash
go test ./...
```

## Installation

Install the CLI:

```bash
go install ./cmd/todo
```

Or run it directly:

```bash
go run ./cmd/todo
```

## Checks

```bash
gofmt -w .
go vet ./...
go test ./...
golangci-lint run
```

## What I Practiced

* Go project structure
* JSON and file operations
* CLI arguments
* Interfaces
* Error handling with `errors.Is`
* `time.Time`
* Unit testing
* Git

This is a small project, but it helped me understand how different parts of a Go application work together.
