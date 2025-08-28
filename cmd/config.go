/*
Copyright © 2025 Hayk Bagdasaryan hb.saryan@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application",
	Long:  `Configure the application`,
	Run: func(cmd *cobra.Command, args []string) {
		configureApp()
	},
}

func configureApp() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	defaultTasksPath := filepath.Join(home, "tasks.json")

	var tasksPath string
	prompt := &survey.Input{
		Message: "Where would you like to store your tasks file?",
		Default: defaultTasksPath,
	}

	err = survey.AskOne(prompt, &tasksPath)
	if err != nil {
		fmt.Println("Error asking for tasks path:", err)
		return
	}

	viper.Set("tasksPath", tasksPath)

	configDir := filepath.Join(home, ".tsk")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Println("Error creating config directory:", err)
		return
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	if err := viper.SafeWriteConfig(); err != nil {
		fmt.Println("Error writing config:", err)
		return
	}

	fmt.Printf("Configuration saved. Tasks will be stored at %s", tasksPath)
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
