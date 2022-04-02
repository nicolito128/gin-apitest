package tasks

import (
	"errors"
	"log"

	"github.com/nicolito128/gin-apitest/pkg/database"
)

// GetTasks() returns all tasks saved in the database
func GetTasks() []Task {
	query := `SELECT * FROM tasks`
	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	list := []Task{}
	for rows.Next() {
		curTask := Task{}
		err = rows.Scan(&curTask.ID, &curTask.Name, &curTask.Content)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, curTask)
	}

	return list
}

// 	CreateTask() insert a new Task in the database
func CreateTask(t Task) error {
	query := `INSERT INTO 
				tasks (name, content)
				VALUES ($1, $2)`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(t.Name, t.Content)
	if err != nil {
		log.Fatal(err)
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("Error: more than one row affected")
	}

	return nil
}

func DeleteTaskById(id int) error {
	query := `DELETE FROM tasks WHERE id = ($1)`

	db := database.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("Error: more than one row affected")
	}

	return nil
}
