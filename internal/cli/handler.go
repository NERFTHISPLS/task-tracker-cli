package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/NERFTHISPLS/task-tracker-cli/internal/task"
)

func Run(service *task.TaskService) {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: task-cli <command> [args]")
		return
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("usage: task-cli add <title> <description>")
			return
		}

		if err := service.Add(args[2]); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("task added successfully")
	case "update":
		if len(args) < 4 {
			fmt.Println("usage: task-cli update <id> <description>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("invalid id, provide integer id")
			return
		}

		if err := service.UpdateDescription(id, args[3]); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("task updated successfully")
	case "delete":
		if len(args) < 3 {
			fmt.Println("usage: task-cli delete <id>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("invalid id, provide integer id")
			return
		}

		if err := service.Delete(id); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("task deleted successfully")
	case "list":
		if len(args) < 3 {
			tasks, err := service.List()
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, t := range tasks {
				fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}

			return
		} else {
			tasks, err := service.ListByStatus(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, t := range tasks {
				fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}
		}
	default:
		fmt.Println("unknown command:", args[1])
	}
}
