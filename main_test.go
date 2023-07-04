package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {
	tasks := []Task{
		{ID: 1, Title: "Task 1", Description: "Example 1", Completed: false},
		{ID: 2, Title: "Task 2", Description: "Example 2", Completed: true},
	}
	addTask(&tasks)

	if len(tasks) != 3 {
		t.Errorf("Expected %d tasks, but got %d", 3, len(tasks))
	}
}
