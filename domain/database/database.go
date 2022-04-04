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

// GetConnection() open de database a return it and a error.
func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Query(query string, args ...any) (*sql.Rows, error) {
	db, err := GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Request() executes multiple requests to the database.
// Prepare the connection (db.Prepare(query)) and then execute the query (db.Exec(stmt).
func Request(query string, args ...any) error {
	db, err := GetConnection()
	if err != nil {
		return err
	}
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
