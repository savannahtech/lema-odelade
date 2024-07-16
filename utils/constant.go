package utils

import (
	"os"
)

const (
	GitHubBaseUrl = "https://github.com"
)

// get data from environment variable
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
