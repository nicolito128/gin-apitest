package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicolito128/tasks-api/domain/queries"
	"github.com/nicolito128/tasks-api/domain/tasks"
)

var TaskList = queries.GetTasks()

func Tasks_GetAllEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, TaskList)
}

func Tasks_FindEndpoint(ctx *gin.Context) {
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

func Tasks_CreateEndpoint(ctx *gin.Context) {
	header := ctx.ContentType()
	if header != "application/json" {
		fmt.Fprintf(ctx.Writer, "Invalid content-type.")
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	newTask := tasks.Task{}
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

	err = queries.CreateTask(newTask)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Task creation failed.")
		return
	}

	TaskList = queries.GetTasks()
	ctx.JSONP(http.StatusOK, newTask)
}

func Tasks_DeleteEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Invalid id.")
		return
	}

	for _, task := range TaskList {
		if task.ID == taskID {
			err = queries.DeleteTaskById(task.ID)
			if err != nil {
				fmt.Fprintf(ctx.Writer, "Task deletion failed.")
				break
			}

			TaskList = queries.GetTasks()
			ctx.String(http.StatusOK, "Task %d deleted succesfully.", taskID)
			break
		}
	}
}

func Tasks_UpdateEndpoint(ctx *gin.Context) {
	id := ctx.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "Invalid id.")
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()

	var newTask tasks.Task
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
			err = queries.UpdateTask(newTask)
			if err != nil {
				fmt.Fprintf(ctx.Writer, "Task update failed.")
				break
			}

			TaskList = queries.GetTasks()
			ctx.String(http.StatusOK, "Task %d updated succesfully!", newTask.ID)
		}
	}
}
