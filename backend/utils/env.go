package utils

import (
	"fmt"
	"os"
)

func GetEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	return env
}

func GetEnvOrDefault(key, defaultValue string) (string, bool) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, false
	}
	return value, true
}

func ValidateEnv(keys []string) {
	for _, key := range keys {
		if _, exists := os.LookupEnv(key); !exists {
			// If the key does not exist or is empty, return false
			fmt.Printf("Environment variable %s is not set\n", key)
			os.Exit(2)
		}
	}
	fmt.Println("All required environment variables are set.")
}
