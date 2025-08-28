/*
Copyright © 2025 Hayk Bagdasaryan hb.saryan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/haykerman/tsk/pkg/config"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Removes all data and starts fresh",
	Long:  `Removes all data and starts fresh`,
	Run: func(cmd *cobra.Command, args []string) {
		err := reset()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

	},
}

func reset() error {
	var filePath, err = config.GetTasksFilePath()
	if err != nil {
		return err
	}

	confirm := false
	prompt := &survey.Confirm{
		Message: "Are you sure you want to reset?",
		Default: false, // change to true if you want default Yes
	}

	err = survey.AskOne(prompt, &confirm)
	if err != nil {
		fmt.Println("Prompt failed:", err)
		return err
	}

	if !confirm {
		fmt.Println("Reset cancelled")
		return nil
	}

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return err
	}

	fmt.Println("Reset completed successfully")
	return nil
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
