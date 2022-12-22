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

package config

import (
	"bufio"
	"fmt"
	"github.com/Zhbert/stats-informer/m/v2/internal/services/env"
	"log"
	"os"
	"time"
)

// GetListOfRepos Function for getting a list of repositories from a configuration file
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
