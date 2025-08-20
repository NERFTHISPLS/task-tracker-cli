package main

import (
	"github.com/NERFTHISPLS/task-tracker-cli/internal/cli"
	"github.com/NERFTHISPLS/task-tracker-cli/internal/task"
	"github.com/NERFTHISPLS/task-tracker-cli/storage"
)

func main() {
	repo := &storage.FileStorage{Path: "tasks.json"}
	service := &task.TaskService{Repo: repo}
	cli.Run(service)
}
