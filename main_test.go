package main

import (
	"bufio"
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestAddWithNonEmptyTitle(t *testing.T) {
	// initialize an empty tasks slice
	var tasks []Task

	// prepare a string with the desired input
	input := "First Task\nTask Description\n"
	reader := strings.NewReader(input)

	// call the addTask func with the test reader
	addTask(&tasks, reader)

	// check if the expected task was added successfully
	expectedTask := Task{
		ID:          1,
		Title:       "First Task",
		Description: "Task Description",
		Completed:   false,
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, but go %d", len(tasks))
	}

	if !reflect.DeepEqual(tasks[0], expectedTask) {
		t.Errorf("Expected task %#v, but got %#v", expectedTask, tasks[0])
	}

}

func TestAddTaskWithEmptyTitle(t *testing.T) {
	tasks := []Task{}

	reader := bufio.NewReader(strings.NewReader(""))

	err := addTask(&tasks, reader)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != "title is required" {
		t.Errorf("Expected error message to be 'title is required', got %s", err.Error())
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}
func TestCompleteTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Task 2", Description: "Description 2", Completed: true},
	}
	reader := bufio.NewReader(strings.NewReader("1"))

	err := completeTask(tasks, reader)

	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if tasks[0].Completed != true {
		t.Errorf("Expected task to be completed, got %v", tasks[0].Completed)
	}

}

func TestCompleteTaskWithInvalidID(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Task 2", Description: "Description 2", Completed: true},
	}
	reader := bufio.NewReader(strings.NewReader("4"))

	err := completeTask(tasks, reader)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if err.Error() != "invalid ID" {
		t.Errorf("Expected error message 'invalid ID', got %v", err.Error())
	}
}

func TestListTasks(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Task 2", Description: "Description 2", Completed: true},
	}

	// Create a buffer to capture the output
	buffer := bytes.Buffer{}

	// Call the listTasks function with the buffer as the writer
	listTasks(tasks, &buffer)

	// Retrieve the printed output from the buffer
	actualOutput := buffer.String()

	// Define the expected output
	expectedOutput := "ID: 1\nTitle: Task 1\nDescription: Description 1\nCompleted: false\n\n" +
		"ID: 2\nTitle: Task 2\nDescription: Description 2\nCompleted: true\n\n"

	// Split the expected and actual output into lines
	expectedLines := strings.Split(expectedOutput, "\n")
	actualLines := strings.Split(actualOutput, "\n")

	// Compare each line of the expected and actual output
	for i, expectedLine := range expectedLines {
		if i >= len(actualLines) {
			t.Errorf("Expected output has more lines than actual output")
			break
		}

		actualLine := actualLines[i]

		if expectedLine != actualLine {
			t.Errorf("Expected line %d:\n%s\n\nBut got:\n%s\n", i+1, expectedLine, actualLine)
		}
	}

	// Check if actual output has more lines than expected output
	if len(actualLines) > len(expectedLines) {
		t.Errorf("Actual output has more lines than expected output")
	}
}
