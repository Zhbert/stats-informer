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
	"github.com/Zhbert/stats-informer/m/v2/internal/common"
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers/structs"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GitHubPage Controller of a page with GitHub statistics
func GitHubPage(context *gin.Context) {
	reposList := config.GetListOfRepos()
	dataOfRepos := make([]structs.ViewData, 0)

	limit, count := 0, 0

	context.HTML(http.StatusOK, "github.tmpl", gin.H{
		"title":      "GitHub Projects",
		"data":       dataOfRepos,
		"respLimit":  limit,
		"respCount":  count,
		"version":    common.GetVersion(),
		"typeOfPage": "full",
		"totalRepos": len(reposList)})
}

// GitHubShortPage Controller of a short page with GitHub statistics
func GitHubShortPage(context *gin.Context) {
	reposList := config.GetListOfRepos()
	dataOfRepos := make([]structs.ViewData, 0)

	limit, count := 0, 0

	context.HTML(http.StatusOK, "github-short.tmpl", gin.H{
		"title":      "GitHub Short Projects",
		"data":       dataOfRepos,
		"respLimit":  limit,
		"respCount":  count,
		"version":    common.GetVersion(),
		"typeOfPage": "short",
		"totalRepos": len(reposList)})
}
