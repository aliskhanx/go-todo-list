package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID     int
	Title  string
	Status string
}
type Tasks map[int]Task

func (t *Tasks) PrintAllTasks() {
	if len(*t) == 0 {
		fmt.Print("\nNo tasks for today. Do you want add a new one?\n")
		return
	}

	fmt.Print("\nHere your tasks:\n\n")
	for k, v := range *t {
		title := strings.TrimSpace(v.Title)
		fmt.Printf("Task %v. %v. Status: %s\n", k, title, v.Status)
	}
}

func (t *Tasks) Add(task Task) {
	(*t)[task.ID] = task
	fmt.Print("\nNew task was added\n")

	err := t.SaveToFile("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (t *Tasks) Edit(task Task) {
	_, ok := (*t)[task.ID]
	if !ok {
		fmt.Println("No task found")
		return
	}

	(*t)[task.ID] = task
	fmt.Printf("\nTask %d was successfully updated to: %s", task.ID, task.Title)

	err := t.SaveToFile("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (t *Tasks) Remove(id *int, taskID int) {
	delete(*t, taskID)
	fmt.Printf("\nTask %v successfully deleted\n", *id)

	*id--
	for i := range *t {
		task := (*t)[i]
		task = Task{ID: task.ID - 1, Title: task.Title, Status: task.Status}
	}

	err := t.SaveToFile("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (t *Tasks) GetMuxID() int {
	maxID := 0
	for i := range *t {
		id := (*t)[i].ID
		if id > maxID {
			maxID = id
		}
	}
	return maxID
}

func (t *Tasks) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(t); err != nil {
		return fmt.Errorf("unable to encode file: %v", err)
	}

	fmt.Println("\n\nTasks saved successfully.")
	return nil
}

func (t *Tasks) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(t); err != nil {
		return fmt.Errorf("unable to load tasks from file: %v", err)
	}

	fmt.Println("\nTasks loaded successfully.")
	return nil
}
