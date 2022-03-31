package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

// Connection vars
var (
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	port     = os.Getenv("DB_PORT")
	database = os.Getenv("DB_NAME")

	connectionString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	Db = db
	fmt.Println("/*************************************/")
	fmt.Println("    Database connected succesfully!    ")
	fmt.Println("/*************************************/")
	return db, nil
}
