package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	tasks := Tasks{}

	err := tasks.LoadFromFile("tasks.json")
	if err != nil {
		fmt.Println("No tasks file found, starting with an empty list.")
	}

	id := tasks.GetMuxID() + 1

	for {
		fmt.Print("\nChoose an option:\n\n" +
			"v - view tasks\n" +
			"c - create task\n" +
			"e - edit task\n" +
			"d - delete task\n" +
			"q - quit\n\n",
		)
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "v":
			tasks.PrintAllTasks()
		case "c":
			fmt.Print("\nWhat do you want to do? ")
			taskTitle, _ := in.ReadString('\n')
			task := Task{ID: id, Title: taskTitle, Status: "pending"}
			tasks.Add(task)
			id++
		case "e":
			var taskId int
			tasks.PrintAllTasks()

			fmt.Print("\nSelect task to edit (enter a number of task): ")
			if _, err := fmt.Scanln(&taskId); err != nil {
				fmt.Println("Invalid input. Try again.")
				continue
			}

			var task Task
			fmt.Printf("\nEditing task %d: ", taskId)
			taskTitle, _ := in.ReadString('\n')

			// Remove the newline character that `ReadString` reads
			// Otherwise if the task status = done the output will be like this
			// Read a book
			// ✓
			taskTitle = taskTitle[:len(taskTitle)-1]

			fmt.Print("\nMark as done? (y/n): ")
			var markDone string
			fmt.Scanln(&markDone)

			if markDone == "y" || markDone == "Y" {
				task = Task{ID: taskId, Title: taskTitle, Status: "completed"}
			} else {
				task = Task{ID: taskId, Title: taskTitle, Status: "pending"}
			}

			tasks.Edit(task)
		case "d":
			var taskId int
			tasks.PrintAllTasks()

			fmt.Print("\nSelect task to delete (enter a number of task): ")
			if _, err := fmt.Scanln(&taskId); err != nil {
				fmt.Println("Invalid input. Try again.")
				continue
			}

			tasks.Remove(&id, taskId)
		case "q":
			tasks.PrintAllTasks()
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
