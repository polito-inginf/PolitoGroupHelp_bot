package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(name string) (string) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(fmt.Sprintf("failed to load env variables file: %v", err))
	}

	result := os.Getenv(name)
	if result == "" {
		panic(fmt.Sprintf("%v env variable is missing", name))
	}

	return result
}
