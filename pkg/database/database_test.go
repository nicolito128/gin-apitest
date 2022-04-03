package database_test

import (
	"testing"

	"github.com/nicolito128/tasks-api/pkg/database"
)

func Test_GetConnection(t *testing.T) {
	_, err := database.GetConnection()
	if err != nil {
		t.Errorf("GetConnection() returned an error: %s", err)
		t.Fail()
	} else {
		t.Log("GetConnection() passed the test.")
	}
}
