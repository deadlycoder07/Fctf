package utils

import (
	"os"

	"github.com/go-pg/pg/v10"
)

func GetPostgresConnection() *pg.DB {
	//read enviroment variables
	url, _ := ConnectionURLBuilder("postgres")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	//connect to database
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     url,
	})

	return db

}
