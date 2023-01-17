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

package app

import (
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Run The function that starts the web server
func Run() {
	route := gin.Default()

	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "main.tmpl", gin.H{"title": "Stats Informer"})
	})
	route.GET("/github", controllers.GitHubPage)
	route.POST("/github", controllers.GetGitHubPage)

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
