package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	client *sql.DB
}

var DB Database

func Connect() {
	var connString string = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")
	var err error
	DB.client, err = sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
}

func (db Database) Close() {
	db.client.Close()
}
