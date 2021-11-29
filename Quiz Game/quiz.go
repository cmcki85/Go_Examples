package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func formatProblems(data [][]string, timeLimit int) {
	var solution string
	var scoredPoints int
	totalPoints := len(data)

	gameTimer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i := range data {
		fmt.Println("What is " + data[i][0] + "?")
		answerChannel := make(chan string)
		go func() {
			fmt.Scan(&solution)
			answerChannel <- solution
		}()

		select {
		case <-gameTimer.C:
			fmt.Printf("\nYou scored: %d, out of %d!", scoredPoints, totalPoints)
			return
		case solution := <-answerChannel:
			if solution == data[i][1] {
				scoredPoints += 1
			}
		}
	}
	fmt.Printf("You scored: %d, out of %d!", scoredPoints, totalPoints)
}

func main() {
	fileName := flag.String("fp", "./problems.csv", "The default filepath to parsable CSV.")
	timeLimit := flag.Int("limit", 30, "The default time limit of the game. (seconds)")
	flag.Parse()

	fp, err := os.Open(*fileName)
	if err != nil {
		log.Fatal("Unable to read file path "+*fileName, err)
	}
	defer fp.Close()

	csvReader := csv.NewReader(fp)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV file for "+*fileName, err)
	}

	formatProblems(data, *timeLimit)
}
