package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

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

	for _, q := range questions {
		fmt.Println(q)
	}
}
