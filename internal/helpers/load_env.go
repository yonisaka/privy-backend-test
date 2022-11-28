package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	result := os.Getenv(key)

	return result
}
