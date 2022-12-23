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

package github

import (
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/structs"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// GetRepoInfo The function of getting data from the repository using the GitHub API
func GetRepoInfo(path string) structs.GitHubStruct {
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := structs.GitHubStruct{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	return data
}

// GetCountsOfResponses Function for getting GitHub API limits for the current session
func GetCountsOfResponses() (int, int) {
	limit := 0
	remaining := 0

	resp, respError := http.Get("https://api.github.com/users/octocat")
	if respError != nil {
		log.Fatalln(respError)
	}
	if i, err := strconv.Atoi(resp.Header.Get("x-ratelimit-limit")); err == nil {
		limit = i
	}
	if i, err := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining")); err == nil {
		remaining = i
	}
	return limit, remaining
}
