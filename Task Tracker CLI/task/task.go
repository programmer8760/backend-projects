package task

import "time"

type Task struct {
	ID          int
	Description string
	Status      status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type status int

const (
	Todo = iota
	InProgress
	Done
)

func Make(id int, description string, status status, createdAt, updatedAt time.Time) Task {
	return Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt}
}

func (s status) String() string {
	switch s {
	case Todo:
		return "todo"
	case InProgress:
		return "in-progress"
	case Done:
		return "done"
	default:
		return "unknown-status"
	}
}
