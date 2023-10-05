package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(name string) (string) {
	result, err := LoadEnvIfPresent(name)
	if err != nil {
		panic(err)
	}

	return result
}

func LoadEnvIfPresent (name string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("failed to load env variables file: %v", err)
	}

	result := os.Getenv(name)
	if result == "" {
		return "", fmt.Errorf("%v env variable is missing", name)
	}

	return result, nil
}

func IsEnvPresent (name string) (bool) {
	_, err := LoadEnvIfPresent(name)
	return err == nil
}
