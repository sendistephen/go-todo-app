package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func main() {
	var tasks []Task // stores a list of tasks
	addTask(&tasks)
	fmt.Println(tasks)
}

// Handles the logic for adding a task to the list
func addTask(tasks *[]Task) {
	var title, description string

	fmt.Println("Enter task title: ")
	fmt.Scanln(&title)

	bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("Enter task description: ")
	description, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	description = strings.TrimSpace(description)

	task := Task{
		ID:          len(*tasks) + 1,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	// append task to tasks
	*tasks = append(*tasks, task)

	fmt.Println("Task added successfully!")
}
