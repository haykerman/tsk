/*
Copyright © 2025 Hayk Bagdasaryan hb.saryan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/haykerman/tsk/pkg/task"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tasks",
	Long:  `List your tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		err := listTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
		}
	},
}

func listTasks() error {
	taskList, err := task.LoadTaskList()
	if err != nil {
		fmt.Println("Error loading task list:", err)
		return err
	}

	fmt.Println("Tasks:")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Title", "Is Completed"})

	for _, task := range taskList.Tasks {
		status := "✗" // Unicode U+2717
		if task.Completed {
			status = "✅"
		}
		table.Append([]string{task.Title, status})
	}

	table.Render()

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
