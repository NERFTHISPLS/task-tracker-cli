package task

import (
	"time"

	"github.com/NERFTHISPLS/task-tracker-cli/internal/e"
)

const (
	addTaskErr = "failed to add task"
)

type TaskService struct {
	repo Repository
}

func (s *TaskService) Add(description string) error {
	tasks, err := s.repo.List()
	if err != nil {
		return e.Wrap(addTaskErr, err)
	}

	id := len(tasks) + 1
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Add(task)
}
