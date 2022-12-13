package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome to the documentation verification service!")
	})

	err := route.Run()
	if err != nil {
		return
	}
}
