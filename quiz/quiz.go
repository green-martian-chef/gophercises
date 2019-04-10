package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a CSV file to be parsed in a format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("The CSV file could not be opened.\n%s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.TrimLeadingSpace = true

	questions, err := r.ReadAll()
	if err != nil {
		fmt.Printf("The CSV file could not be parsed.\n%s\n", err)
		os.Exit(1)
	}

	problems := make([]problem, len(questions))

	for i, p := range questions {
		problems[i] = problem{
			q: p[0],
			a: p[1],
		}
	}

	var correct int
	var answer string

	for i, p := range problems {
		fmt.Printf("Problem %2d : %s = ", i+1, p.q)
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d", correct)
}
