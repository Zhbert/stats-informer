package github

import (
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/structs"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"net/http"
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
