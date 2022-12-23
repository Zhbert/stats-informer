/*
 * Copyright (c) 2022 Konstantin <Zhbert> Nezhbert.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package controllers

import (
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers/structs"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/config"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/github"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GitHubPage Controller of a page with GitHub statistics
func GitHubPage(context *gin.Context) {
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
		"totalRepos": len(reposList)})
}
