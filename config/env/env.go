package env

import (
	"fmt"
	log "github.com/gofiber/fiber/v2/log"
	"os"

	"github.com/joho/godotenv"
)

// sanityCheck checks if all required environment variables are defined
func sanityCheck() {
	requiredEnvVars := []string{
		"API_HOST",
		"API_PORT",
		"HOST",
		"PORT",
		"USER",
		"PASS",
		"SCHEMA",
	}

	for _, envVar := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf("Environment variable %s not defined. Terminating application...", envVar)
		}
	}
}

// LoadConfig loads configuration values from a file using godotenv package
func LoadConfig(filePath string) (*CONFIG, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
		return nil, fmt.Errorf("error loading .env: %v", err)
	}

	sanityCheck()

	return &CONFIG{
		MICRO: MICRO{
			API: API{
				HOST: os.Getenv("API_HOST"),
				PORT: os.Getenv("API_PORT"),
			},
		},
	}, nil
}