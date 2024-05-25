package main

import (
	"log"
	"os"
)

// %%%%%%%%%% Load Data %%%%%%%%%%

func loadBytesFromTextFile(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
