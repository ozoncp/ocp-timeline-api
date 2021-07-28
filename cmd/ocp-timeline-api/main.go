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

func readConfig() error {
	configPath := getPathToConfig()

	for i := 0; i < 5; i++ {
		file, err := os.Open(configPath)

		if err != nil {
			return err
		}

		defer func() {
			file.Close()
		}()
	}

	return nil
}

func getPathToConfig() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Join(dir, configFileName)
}
