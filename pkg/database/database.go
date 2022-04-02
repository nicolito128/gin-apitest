package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var dsn = os.Getenv("DATABASE_URL")

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
