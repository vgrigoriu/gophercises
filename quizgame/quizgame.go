package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type quiz struct {
	question string
	answer   string
}

var right = 0
var quizes []quiz

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
	quizes = make([]quiz, len(records))
	for i, record := range records {
		quizes[i] = quiz{record[0], record[1]}
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("You have 30 seconds, press enter when you're ready to begin...")
	scanner.Scan()
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// start two goroutines and wait for the first of them to finish
	go takeQuiz(scanner, ch1)
	go startTimer(ch2)
	select {
	case <-ch1:
	case <-ch2:
		fmt.Println()
		fmt.Println("Time expired.")
	}

	fmt.Printf("You got %d out of %d correct.", right, len(quizes))
}

func startTimer(signal chan struct{}) {
	time.Sleep(10 * time.Second)
	close(signal)
}

func takeQuiz(scanner *bufio.Scanner, signal chan struct{}) {
	for _, quiz := range quizes {
		fmt.Print(quiz.question)
		fmt.Print(" = ")
		scanner.Scan()
		answer := scanner.Text()
		if answer == quiz.answer {
			right++
		}
	}

	close(signal)
}

func filename() string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "problems.csv"
	}

	return args[0]
}
