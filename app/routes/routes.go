package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nicolito128/tasks-api/app/controllers"
	"github.com/nicolito128/tasks-api/app/middlewares"
)

var router = gin.New()

func SetDefault() *gin.Engine {
	setConfiguration()

	router.GET("/", controllers.IndexEndpoint)

	router.GET("/tasks", controllers.Tasks_GetAllEndpoint)

	router.GET("/tasks/:id", controllers.Tasks_FindEndpoint)

	router.POST("/tasks", middlewares.IsNotJson, middlewares.HeaderOptions, controllers.Tasks_CreateEndpoint)

	router.DELETE("/tasks/:id", middlewares.IsNotJson, middlewares.HeaderOptions, controllers.Tasks_DeleteEndpoint)

	router.PUT("/tasks/:id", middlewares.IsNotJson, middlewares.HeaderOptions, controllers.Tasks_UpdateEndpoint)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}

	return router
}

func setConfiguration() {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	// Files
	router.Static("/static", "./public")
	router.LoadHTMLFiles("./public/index.html")

	// Mode
	mode := os.Getenv("MODE")
	if mode != "" && mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
