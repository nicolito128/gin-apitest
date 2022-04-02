package database

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var ErrRowsAffected = errors.New("Error: more than one row affected")

var dsn = os.Getenv("DATABASE_URL")

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Request(query string, args ...any) error {
	db := GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return ErrRowsAffected
	}

	return nil
}
