package controllers

import (
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers/structs"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GitHubPage(context *gin.Context) {
	data := github.GetRepoInfo("https://api.github.com/repos/werf/werf")
	viewData := structs.ViewData{}

	viewData.Name = data.Name

	context.HTML(http.StatusOK, "github.tmpl", gin.H{
		"title": "GitHub Projects",
		"data":  viewData})
}
