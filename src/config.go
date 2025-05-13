package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	GoogleClientID     string `json:"clientID"`
	GoogleClientSecret string `json:"clientSecret"`
}

// LoadConfig loads the configuration from a file or environment variables
func LoadConfig() (*Config, error) {
	// First try environment variables
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	// If environment variables are set, use them
	if clientID != "" && clientSecret != "" {
		return &Config{
			GoogleClientID:     clientID,
			GoogleClientSecret: clientSecret,
		}, nil
	}
	// Otherwise try to load from credentials.json	// Try current directory first
	data, err := os.ReadFile("credentials.json")
	if err != nil {
		// Try parent directory
		data, err = os.ReadFile("../credentials.json")
		if err != nil {
			return nil, fmt.Errorf("failed to read credentials.json: %v", err)
		}
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse credentials.json: %v", err)
	}

	return &config, nil
}
