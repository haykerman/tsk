/*
Copyright © 2025 Hayk Bagdasaryan hb.saryan@gmail.com
*/
package cmd

import (
	"fmt"

	"time"

	"github.com/google/uuid"
	"github.com/haykerman/tsk/pkg/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the list",
	Long:  `Add a new task to the list`,
	Run: func(cmd *cobra.Command, args []string) {

		// if there are more than one arguments, throw error
		if len(args) > 1 {
			fmt.Println("Error: too many arguments")
			return
		}

		task := args[0]

		err := addTask(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}

		fmt.Println("Task added successfully")
	},
}

func addTask(title string) error {
	fmt.Printf("adding task: %s\n", title)

	taskList, err := task.LoadTaskList()
	if err != nil {
		fmt.Println("Error loading task list:", err)
		return err
	}

	newTask := task.Task{
		ID:        uuid.New().String(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now().String(),
	}
	taskList.Tasks = append(taskList.Tasks, newTask)

	// Save tasklist
	err = task.SaveTaskList(taskList)
	if err != nil {
		fmt.Println("Error saving task list:", err)
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
