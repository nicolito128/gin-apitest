package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nicolito128/gin-apitest/pkg/database"
	"github.com/nicolito128/gin-apitest/pkg/tasks"
)

func main() {
	router := gin.New()
	_, err := database.Init()
	if err != nil {
		log.Fatalf("%s", err)
	}

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies(nil)

	mode := os.Getenv("MODE")
	if mode != "" && mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Files
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./templates/*")

	router.GET("/", indexEndpoint)
	router.GET("/tasks", tasks.GetTasks)
	router.GET("/tasks/:id", tasks.FindTask)
	router.POST("/tasks", tasks.CreateTask)
	router.DELETE("/tasks/:id", tasks.DeleteTask)
	router.PUT("/tasks/:id", tasks.UpdateTask)

	router.Run()
}

func indexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
