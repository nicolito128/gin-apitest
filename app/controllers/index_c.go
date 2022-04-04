package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexEndpoint(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
