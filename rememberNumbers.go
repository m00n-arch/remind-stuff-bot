package main

import (
	"log"
	"os"
)

func creator() {
	csvFile, err := os.Create("remindData.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvFile.Close()
}
