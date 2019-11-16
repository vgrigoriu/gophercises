package main

import (
	"bufio"
	"encoding/csv"
	"flag"
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
	file, err := os.Open(filename())
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

	right := 0
	scanner := bufio.NewScanner(os.Stdin)
	for _, quiz := range quizes {
		fmt.Print(quiz.question)
		fmt.Print(" = ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == quiz.answer {
			right++
		}
	}

	fmt.Printf("got %d out of %d correct", right, len(quizes))
}

func filename() string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "problems.csv"
	}

	return args[0]
}
