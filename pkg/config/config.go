package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// GetTasksFilePath returns the configured path for the tasks file
func GetTasksFilePath() (string, error) {
	// Initialize viper to read the config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}

	configDir := filepath.Join(home, ".tsk")
	viper.AddConfigPath(configDir)

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		return "", fmt.Errorf("error reading config file. Have you run 'tsk config'?: %w", err)
	}

	// Get tasks file path from config
	tasksPath := viper.GetString("tasksPath")
	if tasksPath == "" {
		return "", fmt.Errorf("tasks file path not configured. Please run 'tsk config'")
	}

	return tasksPath, nil
}
