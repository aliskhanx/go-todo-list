package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}
type Tasks map[int]Task

func (t Tasks) PrintAllTasks() {
	if len(t) == 0 {
		fmt.Print("\nNo tasks for today. Do you want add a new one?\n")
		return
	}

	fmt.Print("\nHere your tasks:\n\n")
	for k, v := range t {
		fmt.Printf("%v. %v", k, v.Title)
		if v.Done {
			fmt.Print("✓\n")
		}
	}
}

func (t Tasks) Add(id int, task Task) {
	t[id] = task
	fmt.Print("\nNew task was added\n")

	err := t.SaveToFile("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (t Tasks) Edit(id int, task Task) {
	_, ok := t[id]
	if !ok {
		fmt.Println("No task found")
		return
	}

	t[id] = task
	fmt.Printf("\nTask %d was successfully updated to: %s", id, task.Title)

	err := t.SaveToFile("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (t Tasks) Remove(id int) {
	delete(t, id)
	fmt.Printf("\nTask %v successfully deleted\n", id)

	err := t.SaveToFile("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func (t Tasks) GetMuxID() int {
	maxID := 0
	for id := range t {
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

	fmt.Println("\nTasks saved successfully.")
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
