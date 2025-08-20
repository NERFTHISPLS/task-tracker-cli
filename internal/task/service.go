package task

import (
	"errors"
	"time"

	"github.com/NERFTHISPLS/task-tracker-cli/internal/e"
)

const (
	addTaskErr              = "failed to add task"
	updateTaskErr           = "failed to update task"
	taskEmptyDescriptionErr = "description must not be empty"
	taskStatusInvalidErr    = "status is invalid"
)

const (
	StatusNew        = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
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

func (s *TaskService) Update(id int, description, status string) error {
	target, err := s.repo.ByID(id)
	if err != nil {
		return e.Wrap(updateTaskErr, err)
	}

	if isEmptyString(description) {
		return errors.New(taskEmptyDescriptionErr)
	}

	if !isStatusValid(status) {
		return errors.New(taskStatusInvalidErr)
	}

	target.Description = description
	target.Status = status
	target.UpdatedAt = time.Now()

	return s.repo.Update(target)
}

func (s *TaskService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *TaskService) List() ([]Task, error) {
	return s.repo.List()
}

func (s *TaskService) ListByStatus(status string) ([]Task, error) {
	if !isStatusValid(status) {
		return nil, errors.New(taskStatusInvalidErr)
	}

	return s.repo.ListByStatus(status)
}

func isEmptyString(str string) bool {
	return str == ""
}

func isStatusValid(status string) bool {
	switch status {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	}

	return false
}
