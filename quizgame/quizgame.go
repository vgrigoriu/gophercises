package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading file as csv: %s", err)
	}
	// read lines from csv
	// parse lines into question & response
	// display question
	// read response
	// count right / wrong answers
	// display results
}
