package main

import "fmt"

func AddTask(task string) {
	var tasks []string
	tasks = append(tasks, task)
	fmt.Printf("Tâche ajoutée avec succès : %s\n", task)
}
