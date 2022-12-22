package app

import (
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers"
	"github.com/gin-gonic/gin"
	"log"
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

	resp, rerr := http.Get("https://api.github.com/users/octocat")
	if rerr != nil {
		log.Fatalln(rerr)
	}
	fmt.Println(resp.Header.Get("x-ratelimit-limit"))

	err := route.Run()
	if err != nil {
		return
	}
}
