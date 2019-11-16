package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type quiz struct {
	question string
	answer   string
}

func main() {
	// read file
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading file as csv: %s", err)
	}

	// parse into question & response
	quizes := make([]quiz, len(records))
	for i, record := range records {
		quizes[i] = quiz{record[0], record[1]}
	}

	for _, quiz := range quizes {
		fmt.Println(quiz)
	}
	// display question
	// read response
	// count right / wrong answers
	// display results
}
