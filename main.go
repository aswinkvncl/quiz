package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "file containing questions and its answers")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s", *csvFile))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		exit(fmt.Sprintf("Error during file read"))
	}
	problems := parseCsvFile(lines)
	correctAnsCount := 0
	for i, p := range problems {
		fmt.Printf("Question #%d : %s\n", i, p.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.answer {
			correctAnsCount++
		}
	}
	fmt.Printf("Score: %d out of %d", correctAnsCount, len(problems))
}

func parseCsvFile(lines [][]string) []problem {
	queAndAns := make([]problem, len(lines))
	for i, line := range lines {
		queAndAns[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return queAndAns
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}