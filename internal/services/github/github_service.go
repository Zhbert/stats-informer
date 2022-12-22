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
