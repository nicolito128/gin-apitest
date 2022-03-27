package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var TaskList []Task = []Task{
	{1, "Task 1", "Task description."},
}

func GetTasksEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, TaskList)
}

func PostTaskEndpoint(c *gin.Context) {
	header := c.ContentType()
	if header != "application/json" {
		fmt.Fprintf(c.Writer, "Invalid content-type!")
		return
	}

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()

	var newTask Task
	err := decoder.Decode(&newTask)
	if err != nil {
		fmt.Fprintf(c.Writer, "Decode failed!")
		return
	}

	// Setting task ID
	newTask.ID = len(TaskList) + 1

	// Add to the TaskList
	TaskList = append(TaskList, newTask)
	c.JSONP(http.StatusOK, newTask)
}
