package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const fileName = "tasks.json"

// LoadTasks
func LoadTasks() ([]Task, error) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
// SaveTasks
func SaveTasks(tasks []Task) error {
	fileData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, fileData, 0644)
}
// AddTask
func AddTask(description string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Erreur lors du chargement : %v\n", err)
		return
	}
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	now := time.Now()
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	if err != nil {
		fmt.Printf("Erreur lors de la sauvegarde : %v\n", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newID)
}

// UpdateTaskStatus
func UpdateTaskStatus(id int, newStatus string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Erreur : Tâche avec l'ID %d introuvable.\n", id)
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Printf("Erreur lors de la sauvegarde : %v\n", err)
		return
	}

	fmt.Printf("Task %d marked as %s successfully\n", id, newStatus)
}
