package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var (
	connectionURL = os.Getenv("DATABASE_URL")
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return nil, err
	}

	Db = db
	fmt.Println("/*************************************/")
	fmt.Println("    Database connected succesfully!    ")
	fmt.Println("/*************************************/")
	return db, nil
}
