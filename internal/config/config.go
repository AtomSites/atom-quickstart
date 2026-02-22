// Package config provides configuration loading for the application.
package config

import (
	"os"
	"strconv"
)

// GetEnvOrDefault gets an environment variable with a fallback default value.
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvOrPanic gets an environment variable or panics if not set.
func GetEnvOrPanic(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	panic(key + " must be defined in env")
}

// GetEnvOrDefaultInt gets an environment variable as int with a fallback default value.
func GetEnvOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

// GetEnvOrDefaultBool gets an environment variable as bool with a fallback default value.
func GetEnvOrDefaultBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
