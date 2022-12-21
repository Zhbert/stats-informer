package config

import (
	"bufio"
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/env"
	"log"
	"os"
	"time"
)

func GetListOfRepos() []string {
	listOfRepos := make([]string, 0)
	filename, err := env.GetPathToFile()
	if !err {
		_, fileError := os.Stat(filename)
		if fileError != nil {
			if os.IsNotExist(fileError) {
				fmt.Println(time.Now().Format(time.RFC850) + ": Config file does not exist")
			}
		} else {
			file, openError := os.Open(filename)
			if openError != nil {
				log.Fatalf("Error when opening file: %s", openError)
			}
			fileScanner := bufio.NewScanner(file)
			for fileScanner.Scan() {
				listOfRepos = append(listOfRepos, fileScanner.Text())
			}
			if scanError := fileScanner.Err(); scanError != nil {
				log.Fatalf("Error while reading file: %s", scanError)
			}

			closeError := file.Close()
			if closeError != nil {
				return nil
			}
		}
	}
	return listOfRepos
}
