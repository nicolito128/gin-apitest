package database_test

import (
	"testing"

	"github.com/nicolito128/tasks-api/domain/database"
)

func Test_GetConnection(t *testing.T) {
	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		t.Errorf("GetConnection() returned an error: %s", err)
		t.Fail()
	} else {
		t.Log("GetConnection() passed the test.")
	}
}

func Test_Query(t *testing.T) {
	rows, err := database.Query("SELECT id FROM tasks")
	defer rows.Close()

	if err != nil {
		t.Errorf("Query() returned an error: %s", err)
		t.Fail()
	} else {
		t.Log("Query() passed the test.")
	}
}

func Test_Request(t *testing.T) {
	err := database.Request("SELECT * FROM tasks LIMIT 1")
	if err != nil {
		t.Errorf("Request() returned an error: %s", err)
		t.Fail()
	} else {
		t.Log("Request() passed the test.")
	}
}
