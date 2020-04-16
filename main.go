package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv fil in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, error := os.Open(*csvFileName)

	if error != nil {
		fmt.Println(error)
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, error := r.ReadAll()

	if error != nil {
		fmt.Println(error)
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for index, problem := range problems {
        fmt.Printf("Problem #%d: %s = ", index+1, problem.question)

        answerCh := make(chan string)
        go func() {
            var answer string
        	fmt.Scanf("%s\n", &answer)
            answerCh <- answer
        }()

        select {
        case <-timer.C:
            fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
            return
        case answer := <-answerCh:
            if answer == problem.answer {
                correct++
            }
        }
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {

	fmt.Println(msg)
	os.Exit(1)
}

func isNumeric(s string) bool {

	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
