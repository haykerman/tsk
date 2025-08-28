package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/haykerman/tsk/pkg/config"
)

// Task represents a single task
type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"createdAt"` // Store as string for easier JSON handling
}

// TaskList represents a list of tasks
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

// LoadTaskList loads the task list from the tasks file
func LoadTaskList() (TaskList, error) {
	var taskList TaskList

	tasksPath, err := config.GetTasksFilePath()
	if err != nil {
		return taskList, err
	}

	// Check if file exists
	if _, err := os.Stat(tasksPath); os.IsNotExist(err) {
		// Return empty task list if file doesn't exist
		return taskList, nil
	}

	// File exists, read it
	data, err := os.ReadFile(tasksPath)
	if err != nil {
		return taskList, fmt.Errorf("error reading tasks file: %w", err)
	}

	// Parse JSON
	if err := json.Unmarshal(data, &taskList); err != nil {
		return taskList, fmt.Errorf("error parsing tasks file: %w", err)
	}

	return taskList, nil
}

// SaveTaskList saves the task list to the tasks file
func SaveTaskList(taskList TaskList) error {
	tasksPath, err := config.GetTasksFilePath()
	if err != nil {
		return err
	}

	// Save to file
	jsonData, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %w", err)
	}

	// Ensure the directory exists
	taskDir := filepath.Dir(tasksPath)
	if err := os.MkdirAll(taskDir, 0755); err != nil {
		return fmt.Errorf("error creating tasks directory: %w", err)
	}

	if err := os.WriteFile(tasksPath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing tasks file: %w", err)
	}

	return nil
}
