package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCSV(filepath string) [][]string {
	fp, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read file path "+filepath, err)
	}
	defer fp.Close()

	csvReader := csv.NewReader(fp)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV file for "+filepath, err)
	}

	return data
}

func formatProblems(data [][]string) {
	var solution string
	var scoredPoints int

	for i := range data {
		fmt.Println("What is " + data[i][0] + "?")

		fmt.Scan(&solution)
		if solution == data[i][1] {
			scoredPoints += 1
		}
	}
	totalPoints := len(data)
	fmt.Printf("You scored: %d, out of %d!", scoredPoints, totalPoints)
}

func main() {
	data := readCSV("./problems.csv")
	formatProblems(data)
}
