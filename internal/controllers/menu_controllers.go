package controllers

import (
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers/structs"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/config"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GitHubPage(context *gin.Context) {
	reposList := config.GetListOfRepos()
	for _, url := range reposList {
		fmt.Println(url)
	}

	data := github.GetRepoInfo("https://api.github.com/repos/werf/werf")
	viewData := structs.ViewData{}

	viewData.Name = data.Name
	viewData.FullName = data.FullName
	viewData.Private = data.Private
	viewData.Description = data.Description
	viewData.HomePage = data.Homepage
	viewData.Stars = data.StargazersCount
	viewData.Watchers = data.WatchersCount
	viewData.OpenIssues = data.OpenIssuesCount
	viewData.Forks = data.ForksCount
	viewData.License = data.License.Name
	viewData.GitHubURL = data.HTMLURL

	context.HTML(http.StatusOK, "github.tmpl", gin.H{
		"title": "GitHub Projects",
		"data":  viewData})
}
