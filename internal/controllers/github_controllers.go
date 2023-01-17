package controllers

import (
	"github.com/Zhbert/stats-informer/m/v2/internal/common"
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers/structs"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/config"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetGitHubPage POST Controller of a page with GitHub statistics
func GetGitHubPage(context *gin.Context) {
	reposList := config.GetListOfRepos()
	dataOfRepos := make([]structs.ViewData, 0)
	for _, url := range reposList {
		data := github.GetRepoInfo(url)
		ViewData := structs.ViewData{}

		ViewData.Name = data.Name
		ViewData.FullName = data.FullName
		ViewData.Private = data.Private
		ViewData.Description = data.Description
		ViewData.HomePage = data.Homepage
		ViewData.Stars = data.StargazersCount
		ViewData.Watchers = data.WatchersCount
		ViewData.OpenIssues = data.OpenIssuesCount
		ViewData.Forks = data.ForksCount
		ViewData.License = data.License.Name
		ViewData.GitHubURL = data.HTMLURL

		dataOfRepos = append(dataOfRepos, ViewData)
	}

	limit, count := github.GetCountsOfResponses()

	context.HTML(http.StatusOK, "github.tmpl", gin.H{
		"title":      "GitHub Projects",
		"data":       dataOfRepos,
		"respLimit":  limit,
		"respCount":  count,
		"version":    common.GetVersion(),
		"totalRepos": len(reposList)})
}
