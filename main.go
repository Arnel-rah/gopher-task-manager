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
			fmt.Println("Erreur: Veuillez fournir une description pour la tâche.")
			return
		}
		AddTask(os.Args[2])

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Erreur : ID manquant")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		UpdateTaskStatus(id, "in-progress")
		
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Erreur : ID manquant")
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
		fmt.Println("Commande inconnue")
	}
}
