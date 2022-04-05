package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicolito128/tasks-api/app/controllers"
)

var router = gin.New()

func Run() *gin.Engine {
	setConfiguration()

	router.GET("/", controllers.IndexEndpoint)

	router.GET("/tasks", headerOptions, controllers.Tasks_GetAllEndpoint)

	router.GET("/tasks/:id", isNotJson, headerOptions, controllers.Tasks_FindEndpoint)

	router.POST("/tasks", isNotJson, headerOptions, controllers.Tasks_CreateEndpoint)

	router.DELETE("/tasks/:id", isNotJson, headerOptions, controllers.Tasks_DeleteEndpoint)

	router.PUT("/tasks/:id", isNotJson, headerOptions, controllers.Tasks_UpdateEndpoint)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}

	return router
}
