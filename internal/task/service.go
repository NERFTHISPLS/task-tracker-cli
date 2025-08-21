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

type Repository interface {
	Add(task Task) error
	Update(task Task) error
	Delete(id int) error
	List() ([]Task, error)
	ListByStatus(status string) ([]Task, error)
	ByID(id int) (Task, error)
}

type TaskService struct {
	Repo Repository
}

func (s *TaskService) Add(description string) error {
	tasks, err := s.Repo.List()
	if err != nil {
		return e.Wrap(addTaskErr, err)
	}

	id := tasks[len(tasks)-1].ID + 1
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.Repo.Add(task)
}

func (s *TaskService) UpdateDescription(id int, description string) error {
	target, err := s.Repo.ByID(id)
	if err != nil {
		return e.Wrap(updateTaskErr, err)
	}

	if isEmptyString(description) {
		return errors.New(taskEmptyDescriptionErr)
	}

	target.Description = description
	target.UpdatedAt = time.Now()

	return s.Repo.Update(target)
}

func (s *TaskService) UpdateStatus(id int, status string) error {
	target, err := s.Repo.ByID(id)
	if err != nil {
		return e.Wrap(updateTaskErr, err)
	}

	if !isStatusValid(status) {
		return errors.New(taskStatusInvalidErr)
	}

	target.Status = status
	target.UpdatedAt = time.Now()

	return s.Repo.Update(target)
}

func (s *TaskService) Delete(id int) error {
	return s.Repo.Delete(id)
}

func (s *TaskService) List() ([]Task, error) {
	return s.Repo.List()
}

func (s *TaskService) ListByStatus(status string) ([]Task, error) {
	if !isStatusValid(status) {
		return nil, errors.New(taskStatusInvalidErr)
	}

	return s.Repo.ListByStatus(status)
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
