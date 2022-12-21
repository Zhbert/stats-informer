package env

import (
	"fmt"
	"os"
	"time"
)

func GetPathToFile() (string, bool) {
	path, pathExists := os.LookupEnv("ST_REPOS_FILE")
	if !pathExists {
		fmt.Println(time.Now().Format(time.RFC850) + ": $ST_REPOS_FILE not found!")
		return "", true
	}
	fmt.Println(time.Now().Format(time.RFC850) + ": Config file: " + path)
	return path, false
}
