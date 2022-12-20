package app

import (
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "main.tmpl", gin.H{"title": "Stats Informer"})
	})
	route.GET("/github", controllers.GitHubPage)

	route.Static("/css", "static/css")
	route.Static("/js", "static/js")
	route.Static("/assets", "static/assets")

	route.LoadHTMLGlob("templates/**/*")

	err := route.Run()
	if err != nil {
		return
	}
}
