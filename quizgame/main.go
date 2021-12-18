package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var answer string
	var correctCount int
	answerCh := make(chan string)

	csvFileName := flag.String("csv", "problems.csv", "csv selects a csv file to use in the quiz game")
	timeLimit := flag.Int("limit", 30, "limit sets the time limit to finish the quiz")
	flag.Parse()

	csvFile, err := os.Open("quizgame/" + *csvFileName)
	if err != nil {
		log.Fatal(err)
	}

	read := csv.NewReader(csvFile)
	records, err := read.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, v := range records {
		fmt.Printf("Question: %s? Your answer:\n", v[0])

		go func() {
			fmt.Scan(&answer)
			answerCh <- answer
		}()
		
		select {
		case <-timer.C:
			fmt.Println("Correct answers:", correctCount)
			return
		case a := <- answerCh:
			if a == v[1] {
				correctCount += 1
			}
			if correctCount == len(records) {
				fmt.Println("Correct answers:", correctCount)
				return
			}
		}
	}
}
