package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"encoding/csv"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv fil in the format of 'question,answer'")
	flag.Parse()
	
	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println(err)
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		exit("Failed to parse the provided CSV file.")
	}
	
	problems := parseLines(lines)
	fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem {
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}