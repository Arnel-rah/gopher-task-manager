package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [command] [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Erreur: Veuillez fournir une description.")
			return
		}
		AddTask(os.Args[2])

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update [id] [description]")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateTaskDescription(id, os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Erreur: ID manquant.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		DeleteTask(id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Erreur: ID manquant.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateTaskStatus(id, "in-progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Erreur: ID manquant.")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateTaskStatus(id, "done")

	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		ListTasks(filter)

	default:
		fmt.Printf("Commande inconnue: %s\n", command)
	}
}

