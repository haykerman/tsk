package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
)

// Initialize sets up the config system
func Initialize(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error getting home directory: %w", err)
		}

		viper.AddConfigPath(filepath.Join(home, ".tsk"))
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

// GetTasksFilePath returns the configured path for the tasks file
func GetTasksFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return "", err
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
		return "", err
	}

	return tasksPath, nil
}
