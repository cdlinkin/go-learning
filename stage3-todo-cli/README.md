# Todo CLI

A small command-line todo app written in Go.

I built this project to practice working with Go project structure, JSON file storage, command-line arguments, error handling, and interfaces.

## What it can do

* Add new tasks
* Show all tasks
* Mark tasks as completed
* Delete tasks
* Store tasks in a JSON file
* Use a custom file path with the `-file` flag

## Project structure

```text
stage3-todo-cli/
├── cmd/
│   └── todo/
│       └── main.go
├── internal/
│   ├── task/
│   │   ├── task.go
│   │   └── errors.go
│   └── storage/
│       ├── storage.go
│       └── file.go
├── README.md
└── go.mod
```

The project is split into a few simple parts:

* `cmd/todo` — CLI entry point and command handling
* `internal/task` — task model and task operations
* `internal/storage` — storage interface and JSON file implementation

## Usage

### Add a task

```bash
todo add "buy some bread"
```

### List tasks

```bash
todo list
```

Example:

```text
1. buy some bread [ ]
2. learn Go [x]
3. finish the project [ ]
```

### Mark a task as done

```bash
todo done 2
```

### Delete a task

```bash
todo delete 3
```

## Custom file

By default, the app stores tasks in `tasks.json` inside the user's home directory.

You can also specify your own file:

```bash
todo -file ./tasks.json add "buy some milk"
```

Then use the same file when listing tasks:

```bash
todo -file ./tasks.json list
```

## Error handling

The project uses a sentinel error for cases when a task cannot be found:

```go
task.ErrTaskNotFound
```

It is checked with `errors.Is`, so the caller can distinguish a missing task from other errors.

File and JSON errors are wrapped with additional context using `%w`.

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd stage3-todo-cli
```

Install the CLI:

```bash
go install ./cmd/todo
```

After that, the `todo` command should be available from your terminal if your Go bin directory is in `PATH`.

## Development

Format the code:

```bash
gofmt -w .
```

Run `go vet`:

```bash
go vet ./...
```

Run the linter:

```bash
golangci-lint run
```

## What's next

This is a basic version of the app. Some possible next steps are:

* ~~Add task priorities~~
* Add deadlines
* ~~Add a `--done` filter for `todo list`~~
* Add tests
* Improve the CLI error messages
