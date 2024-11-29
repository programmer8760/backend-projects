package task

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Status int

const (
	Todo = iota
	InProgress
	Done
)

func Make(id int, description string, status Status, createdAt, updatedAt time.Time) Task {
	return Task{
		ID:          id,
		Description: description,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt}
}

func (s Status) String() string {
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

func (t Task) String() string {
	return fmt.Sprintf("%v\t%s\t%s\t%s\t%s", t.ID, t.Description, t.Status, t.CreatedAt.Format("01.02.06 15:04"), t.UpdatedAt.Format("01.02.06 15:04"))
}

func ParseStatus(s string) Status {
	switch s {
	case "todo":
		return 0
	case "in-progress":
		return 1
	case "done":
		return 2
	default:
		return -1
	}
}

func ChangeStatus(t Task, s Status) Task {
	return Make(t.ID, t.Description, s, t.CreatedAt, t.UpdatedAt)
}
