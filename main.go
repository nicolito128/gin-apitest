package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolito128/gin-apitest/tasks"
)

func main() {
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies(nil)
	gin.SetMode(gin.DebugMode)

	// Files
	router.LoadHTMLGlob("templates/*")
	router.Static("/static/styles", "./static/styles")

	router.GET("/", indexEndpoint)
	router.GET("/tasks", tasks.GetTasks)
	router.GET("/tasks/:id", tasks.FindTask)
	router.POST("/tasks", tasks.CreateTask)
	router.DELETE("/tasks/:id", tasks.DeleteTask)

	router.Run()
}

func indexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
