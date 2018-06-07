package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
	"bufio"
	"time"
)

type CSVRow struct {
	Question string
	Answer string
}

var input string

func main() {
	filePath := flag.String("path", "problems.csv", "path to csv file")
	timeLimit := flag.Int("limit", 5, "time limit per question")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil{
		fmt.Println(err.Error())
	}
	defer f.Close()

	reader := csv.NewReader(bufio.NewReader(f))
	records, err := reader.ReadAll()
	if err != nil{
		fmt.Println(err.Error())
	}

	correctAnswers := 0
	totalQuestions := len(records)
	fmt.Println("Press enter to see next question:")
	fmt.Scanln(&input)

	go func() {
		for _, rec := range  records {
			row := CSVRow{
				Question: rec[0],
				Answer: rec[1],
			}

			fmt.Println(row.Question)
			fmt.Println("Write your answer:")
			fmt.Scanln(&input)
			if input == row.Answer {
				correctAnswers += 1
			}
		}
	}()

	select {
	case <-time.After(time.Duration(int(*timeLimit)) * time.Second):
		fmt.Println("Time is out")
	}

	fmt.Printf("Result: %d/%d", correctAnswers, totalQuestions)
}
