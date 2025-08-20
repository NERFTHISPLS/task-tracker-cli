package storage

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/NERFTHISPLS/task-tracker-cli/internal/e"
	"github.com/NERFTHISPLS/task-tracker-cli/internal/task"
)

const (
	writeFileErr      = "failed to write to file"
	readFileErr       = "failed to read file"
	jsonUnmarshallErr = "failed to parse json"
	jsonMarshallErr   = "failed to encode to json"
	taskNotFoundErr   = "failed to find task"
	tasksEmptyErr     = "tasks list is empty"
)

type FileStorage struct {
	Path string
}

func (fs *FileStorage) Add(task task.Task) error {
	tasks, err := fs.load()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)

	return fs.save(tasks)
}

func (fs *FileStorage) Update(task task.Task) error {
	tasks, err := fs.load()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID != task.ID {
			continue
		}

		tasks[i] = task
		return fs.save(tasks)
	}

	return errors.New(taskNotFoundErr)
}

func (fs *FileStorage) Delete(id int) error {
	tasks, err := fs.load()
	if err != nil {
		return err
	}

	newTasks := make([]task.Task, 0, len(tasks))
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			break
		}

		newTasks = append(newTasks, t)
	}

	if !found {
		return errors.New(taskNotFoundErr)
	}

	return fs.save(newTasks)
}

func (fs *FileStorage) List() ([]task.Task, error) {
	return fs.load()
}

func (fs *FileStorage) ListByStatus(status string) ([]task.Task, error) {
	tasks, err := fs.load()
	if err != nil {
		return nil, err
	}

	newTasks := []task.Task{}

	for _, t := range tasks {
		if t.Status != status {
			continue
		}

		newTasks = append(newTasks, t)
	}

	if len(newTasks) == 0 {
		return nil, errors.New(tasksEmptyErr)
	}

	return newTasks, nil
}

func (fs *FileStorage) ByID(id int) (task.Task, error) {
	tasks, err := fs.load()
	if err != nil {
		return task.Task{}, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return task.Task{}, errors.New(taskNotFoundErr)
}

func (fs *FileStorage) load() ([]task.Task, error) {
	if _, err := os.Stat(fs.Path); os.IsNotExist(err) {
		empty := []task.Task{}
		data, _ := json.Marshal(empty)
		if err := os.WriteFile(fs.Path, data, 0644); err != nil {
			return nil, e.Wrap(writeFileErr, err)
		}

		return empty, nil
	}

	data, err := os.ReadFile(fs.Path)
	if err != nil {
		return nil, e.Wrap(readFileErr, err)
	}

	if len(data) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, e.Wrap(jsonUnmarshallErr, err)
	}

	return tasks, nil
}

func (fs *FileStorage) save(tasks []task.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return e.Wrap(jsonMarshallErr, err)
	}

	if err := os.WriteFile(fs.Path, data, 0644); err != nil {
		return e.Wrap(writeFileErr, err)
	}

	return nil
}
