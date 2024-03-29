package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func main() {
	var tasks []Task // stores a list of tasks

	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addTask(&tasks, os.Stdin)
		case 2:
			listTasks(tasks, os.Stdout)
		case 3:
			completeTask(tasks, os.Stdin)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println()
	}

}

// Handles the logic for adding a task to the list
func addTask(tasks *[]Task, reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	var title string

	for i := 0; i < 3; i++ {
		fmt.Println("Enter task title: ")
		scanner.Scan()
		title = scanner.Text()

		if title != "" {
			break
		}
	}

	if title == "" {
		return errors.New("title is required")
	}

	fmt.Println("Enter task description: ")
	scanner.Scan()
	description := scanner.Text()

	task := Task{
		ID:          len(*tasks) + 1,
		Title:       title,
		Description: description,
		Completed:   false,
	}
	// append task to tasks
	*tasks = append(*tasks, task)

	fmt.Println("Task added successfully!")

	return nil
}

func completeTask(tasks []Task, reader io.Reader) error {

	scanner := bufio.NewScanner(reader)

	fmt.Println("Enter task id: ")
	scanner.Scan()
	taskId := scanner.Text()

	// convert taskId to an int
	id, err := strconv.Atoi(taskId)

	if err != nil {
		fmt.Println("ID must be a number")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			fmt.Println("Task completed successfully")
			return nil
		}

	}
	return errors.New("invalid ID")
}

func listTasks(tasks []Task, w io.Writer) {
	for _, task := range tasks {
		fmt.Fprintf(w, "ID: %d\n", task.ID)
		fmt.Fprintf(w, "Title: %s\n", task.Title)
		fmt.Fprintf(w, "Description: %s\n", task.Description)
		fmt.Fprintf(w, "Completed: %t\n", task.Completed)
		fmt.Fprintln(w)
	}
}
