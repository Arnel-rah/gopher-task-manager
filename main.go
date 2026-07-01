package main

import (
	"fmt"
	"os"
	"strconv"
)

type commandFunc func(args []string) error

var commands = map[string]commandFunc{
	"add":              cmdAdd,
	"update":           cmdUpdate,
	"delete":           cmdDelete,
	"mark-in-progress": cmdMarkInProgress,
	"mark-done":        cmdMarkDone,
	"list":             cmdList,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [command] [args]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	fn, ok := commands[command]
	if !ok {
		fmt.Printf("Commande inconnue: %s\n", command)
		return
	}

	if err := fn(args); err != nil {
		fmt.Println(err)
	}
}

func cmdAdd(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Erreur: Veuillez fournir une description.")
	}
	AddTask(args[0])
	return nil
}

func cmdUpdate(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: task-cli update [id] [description]")
	}
	id, err := parseID(args[0])
	if err != nil {
		return err
	}
	UpdateTaskDescription(id, args[1])
	return nil
}

func cmdDelete(args []string) error {
	id, err := requireID(args)
	if err != nil {
		return err
	}
	DeleteTask(id)
	return nil
}

func cmdMarkInProgress(args []string) error {
	id, err := requireID(args)
	if err != nil {
		return err
	}
	UpdateTaskStatus(id, "in-progress")
	return nil
}

func cmdMarkDone(args []string) error {
	id, err := requireID(args)
	if err != nil {
		return err
	}
	UpdateTaskStatus(id, "done")
	return nil
}

func cmdList(args []string) error {
	filter := ""
	if len(args) > 0 {
		filter = args[0]
	}
	ListTasks(filter)
	return nil
}

func requireID(args []string) (int, error) {
	if len(args) < 1 {
		return 0, fmt.Errorf("Erreur: ID manquant.")
	}
	return parseID(args[0])
}

func parseID(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Erreur: ID invalide (%s)", s)
	}
	return id, nil
}
