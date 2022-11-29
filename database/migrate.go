package main

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pressly/goose"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if strings.Contains(path, "database") {
		path = path[:len(path)-8]
	}

	err = godotenv.Load(string(path) + "/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	var argsRaw = os.Args
	var action = argsRaw[1:]
	if len(action) == 0 {
		log.Fatal("Please provide an action")
	}
	var dir = string(path) + "/database"

	var args = [3]string{}
	args[0] = os.Getenv("DB_CONNECTION")
	args[1] = os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(localhost:" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?&parseTime=true"
	args[2] = action[0]

	driver, dbstring, command := args[0], args[1], args[2]

	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	}

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, dir, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}
