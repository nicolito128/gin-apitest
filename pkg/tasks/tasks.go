package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

var TaskList = GetTasks()

func GetAllEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, TaskList)
}

func FindEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Invalid id.")
		return
	}

	for _, task := range TaskList {
		if task.ID == taskID {
			ctx.JSON(http.StatusOK, task)
			break
		}
	}
}

func CreateEndpoint(ctx *gin.Context) {
	header := ctx.ContentType()
	if header != "application/json" {
		fmt.Fprintf(ctx.Writer, "Invalid content-type.")
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	var newTask Task
	err := decoder.Decode(&newTask)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Decode failed.")
		return
	}
	newTask.ID = len(TaskList) + 1

	if newTask.Name == "" {
		fmt.Fprintf(ctx.Writer, "Task name invalid: empty place.")
		return
	}

	err = CreateTask(newTask)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Task creation failed.")
		return
	}

	TaskList = GetTasks()
	ctx.JSONP(http.StatusOK, newTask)
}

func DeleteEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Invalid id.")
		return
	}

	for _, task := range TaskList {
		if task.ID == taskID {
			err = DeleteTaskById(task.ID)
			if err != nil {
				fmt.Fprintf(ctx.Writer, "Task deletion failed.")
				break
			}

			TaskList = GetTasks()
			ctx.String(http.StatusOK, "Task %d deleted succesfully.", taskID)
			break
		}
	}
}

func UpdateEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Invalid id.")
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	var newTask Task
	err = decoder.Decode(&newTask)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Decode failed!")
		return
	}

	if newTask.Name == "" {
		fmt.Fprintf(ctx.Writer, "Task name/content invalid: empty place.")
		return
	}

	for _, task := range TaskList {
		if task.ID == taskID {
			newTask.ID = taskID
			err = UpdateTask(newTask)
			if err != nil {
				fmt.Fprintf(ctx.Writer, "Task update failed.")
				break
			}

			TaskList = GetTasks()
			ctx.String(http.StatusOK, "Task %d updated succesfully!", newTask.ID)
		}
	}
}
