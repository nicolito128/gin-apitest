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
	router.GET("/tasks", tasks.GetTasksEndpoint)
	router.POST("/tasks", tasks.PostTaskEndpoint)

	// Start server
	router.Run()
}

func indexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
