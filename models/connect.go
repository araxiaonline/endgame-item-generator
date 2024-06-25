package models

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	client *sqlx.DB
}

var DB Database

func Connect() {
	var connString string = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME")
	var err error
	DB.client, err = sqlx.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
}

func (db Database) Close() {
	db.client.Close()
}
