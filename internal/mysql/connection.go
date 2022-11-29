package mysql

import (
	"log"
	"privy-backend-test/internal/helpers"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func Connection() *sqlx.DB {
	conn := helpers.GoDotEnvVariable("DB_CONNECTION")
	host := helpers.GoDotEnvVariable("DB_HOST")
	port := helpers.GoDotEnvVariable("DB_PORT")
	user := helpers.GoDotEnvVariable("DB_USERNAME")
	pass := helpers.GoDotEnvVariable("DB_PASSWORD")
	name := helpers.GoDotEnvVariable("DB_DATABASE")

	if db != nil {
		return db
	}

	var err error

	db, err = sqlx.Open(conn, user+":"+pass+"@tcp("+host+":"+port+")/"+name+"?parseTime=true")
	if err != nil {
		log.Fatalf("could not open mysql db connection: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed ping mysql db connection: %v\n", err)
	}

	// this configuration refer to https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
