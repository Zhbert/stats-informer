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
	v1 "github.com/Zhbert/stats-informer/m/v2/internal/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

// Run The function that starts the web server
func Run() {
	route := gin.Default()

	route.GET("/api/v1/get-version", v1.AppVersion)

	err := route.Run()
	if err != nil {
		return
	}
}
