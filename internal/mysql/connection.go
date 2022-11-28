package mysql

import (
	"fmt"
	"privy-backend-test/internal/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() (db *gorm.DB) {
	host := helpers.GoDotEnvVariable("DB_HOST")
	port := helpers.GoDotEnvVariable("DB_PORT")
	user := helpers.GoDotEnvVariable("DB_USERNAME")
	pass := helpers.GoDotEnvVariable("DB_PASSWORD")
	name := helpers.GoDotEnvVariable("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
