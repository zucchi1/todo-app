package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	tasks, err := loadTasks()
	fmt.Println(err)
	NewTask := Task{ID: nextID(tasks),Title: title, Done: false}
	tasks = append(tasks, NewTask)
	saveTasks(tasks)
}

func ListTasks() {
	tasks, err := loadTasks()
	if tasks == nil {
		fmt.Println(err)
		return
	}else{
		for i := 0; i < len(tasks); i++ {
			status := "[]"
			if tasks[i].Done {
				status = "[x]"
			}
			fmt.Println(tasks[i].ID, ":", tasks[i].Title, status)
		}
	}
}

func CompleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}

	fmt.Printf("Task %d marked as complete\n", id)
}

func DeleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	found := false
	newTasks := []Task{}
	for _, task := range tasks {
		if task.ID == id {
			found = true
		} else {
			newTasks = append(newTasks, task)
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	if err := saveTasks(newTasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}

	fmt.Printf("Task %d deleted\n", id)
}
