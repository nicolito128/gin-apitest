package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	// Start server
	router.Run(":8080")
}

func indexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome! - Go Test",
	})
}
