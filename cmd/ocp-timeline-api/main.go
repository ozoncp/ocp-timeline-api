package main

import (
	"fmt"
	"os"
	"path"
)

const configFileName = "config2.json"

func main() {
	readConfig()
	fmt.Println("In this project we probably will calculate timeline of something 🤔")
}

func readConfig() {
	configPath := getPathToConfig()

	for i := 0; i < 5; i++ {
		file, err := os.Open(configPath)

		defer func() {
			if err == nil {
				file.Close()
			}
		}()

	}
}

func getPathToConfig() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, configFileName)
}
