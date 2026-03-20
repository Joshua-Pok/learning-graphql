package task

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID
	Title  string
	Status string
}
