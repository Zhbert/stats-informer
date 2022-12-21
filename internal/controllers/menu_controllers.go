package controllers

import (
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GitHubPage(context *gin.Context) {
	data := github.GetRepoInfo("https://api.github.com/repos/werf/werf")
	fmt.Println(data)

	context.HTML(http.StatusOK, "github.tmpl", gin.H{"title": "GitHub Projects"})
}
