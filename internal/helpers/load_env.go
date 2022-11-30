package helpers

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// if contains tests
	if strings.Contains(path, "repositories") {
		path = path[:len(path)-21]
	}
	if strings.Contains(path, "usecases") {
		path = path[:len(path)-17]
	}

	// load .env file
	err = godotenv.Load(string(path) + "/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	result := os.Getenv(key)

	return result
}
