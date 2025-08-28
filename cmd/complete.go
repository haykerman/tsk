/*
Copyright © 2025 Hayk Bagdasaryan hb.saryan@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  `Mark a task as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")

		err := complete()
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func complete() error {
	fmt.Println("complete called")
	return nil
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
