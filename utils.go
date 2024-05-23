package main

import (
	"log"
	"os"
)

// %%%%%%%%%% Load Data %%%%%%%%%%

func loadFromTextFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
