package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func sanityCheck() {
	envProps := []string{
		"API_HOST",
		"API_PORT",
	}

	for _, v := range envProps {
		if os.Getenv(v) == "" {
			log.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", v))
		}
	}
}

func LoadConfig(filePath string) (*CONFIG, error) {
	err := godotenv.Load(filePath)
	sanityCheck()
	if err != nil {
		log.Fatal("Erorr loading .env")
		return nil, fmt.Errorf("error loading .env: %v", err)
	}

	return &CONFIG{
		MICRO: MICRO{
			API: API{
				HOST: os.Getenv("API_HOST"),
				PORT: os.Getenv("API_PORT"),
			},
		},
	}, nil
}