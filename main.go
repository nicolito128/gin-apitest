package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	gin.SetMode(gin.DebugMode)
	router.LoadHTMLGlob("templates/*")

	router.GET("/", useTemplate("index"))
	router.GET("/about", useTemplate("about"))
	router.GET("/contact", useTemplate("contact"))

	router.Run()
}

func useTemplate(name string) func(c *gin.Context) {
	file := name + ".html"

	return func(c *gin.Context) {
		c.HTML(http.StatusOK, file, gin.H{
			"title": name,
		})
	}

}
