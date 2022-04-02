package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nicolito128/gin-apitest/pkg/tasks"
)

func main() {
	router := gin.New()

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
	router.GET("/tasks", tasks.GetAllEndpoint)
	router.GET("/tasks/:id", tasks.FindEndpoint)
	router.POST("/tasks", tasks.CreateEndpoint)
	router.DELETE("/tasks/:id", tasks.DeleteEndpoint)
	router.PUT("/tasks/:id", tasks.UpdateEndpoint)

	router.Run()
}

func indexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
