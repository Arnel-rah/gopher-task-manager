package main

import (
	"fmt"
	"os"
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

	case "list":
		fmt.Println("Liste des tâches (bientôt disponible...)")

	default:
		fmt.Println("Commande inconnue")
	}
}
