package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsNotJson(ctx *gin.Context) {
	header := ctx.ContentType()
	fmt.Println(header)
	if header != "application/json" {
		ctx.Status(http.StatusNotAcceptable)
		fmt.Fprintf(ctx.Writer, "Invalid content-type: application not json.")
		ctx.Abort()
		return
	}
}

func HeaderOptions(ctx *gin.Context) {
	contentOptions := ctx.Request.Header.Get("X-Content-Type-Options")
	if contentOptions != "nosniff" {
		ctx.Status(http.StatusNotAcceptable)
		fmt.Fprintf(ctx.Writer, "Invalid: X-Content-Type-Options different from nosniff.")
		ctx.Abort()
		return
	}

	frameOptions := ctx.Request.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		ctx.Status(http.StatusAccepted)
		fmt.Fprintf(ctx.Writer, "Invalid: X-Frame-Options different from deny.")
		ctx.Abort()
		return
	}
}
