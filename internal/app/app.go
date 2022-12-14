package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "layouts/index.tmpl", gin.H{"title": "Stats Informer"})
	})

	route.Static("/css", "static/css")
	route.Static("/js", "static/js")
	route.Static("/assets", "static/assets")

	route.LoadHTMLGlob("templates/**/*")

	err := route.Run()
	if err != nil {
		return
	}
}
