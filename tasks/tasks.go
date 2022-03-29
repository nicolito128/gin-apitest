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

func GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, TaskList)
}

func CreateTask(ctx *gin.Context) {
	header := ctx.ContentType()
	if header != "application/json" {
		fmt.Fprintf(ctx.Writer, "Invalid content-type!")
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	var newTask Task
	err := decoder.Decode(&newTask)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Decode failed!")
		return
	}

	// Setting task ID
	newTask.ID = len(TaskList) + 1

	// Add to the TaskList
	TaskList = append(TaskList, newTask)
	ctx.JSONP(http.StatusOK, newTask)
}
