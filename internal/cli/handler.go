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
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("task added successfully")
	case "update":
		if len(args) < 5 {
			fmt.Println("usage: task-cli update <id> <description> <status>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("error: invalid id, provide integer id")
			return
		}

		if err := service.Update(id, args[3], args[4]); err != nil {
			fmt.Println("error:", err)
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
			fmt.Println("error: invalid id, provide integer id")
			return
		}

		if err := service.Delete(id); err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("task deleted successfully")
	case "list":
		if len(args) < 3 {
			tasks, _ := service.List()

			for _, t := range tasks {
				fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}

			return
		} else {
			tasks, _ := service.ListByStatus(os.Args[2])

			for _, t := range tasks {
				fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, t.Status)
			}
		}
	default:
		fmt.Println("unknown command:", args[1])
	}
}
