package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a CSV file to be parsed in a format 'question,answer'")
	timeLimit := flag.Int("limit", 10, "time limit for the quiz in seconds")
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

	// Here we put the slice returned by ReadAll() into a struct. Why? If instead
	// of a CSV file we had a JSON, YAML or something, we only need to change the
	// parser code, the rest remains the same.
	for i, p := range questions {
		problems[i] = problem{
			q: p[0],
			a: p[1],
		}
	}

	var correct int
	var answer string
	aChannel := make(chan string)

	// NewTimer() creates a new type Timer that will send the current time on its
	// channel after the duration
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// The problem here is you want to ask the user a question while the timer is
	// running but you don't want to accept any answer after the timer has
	// expired. Since NewTimer() returns the time through a channel, you can use
	// select and break the loop when the timer is expired. Select waits on
	// multiple channel operations so you have to put the code that receives the
	// user answer on a goroutine, so if the timer is expired select breaks the
	// loop instead of accept the answer.
loop:
	for i, p := range problems {
		fmt.Printf("Problem %2d : %s = ", i+1, p.q)

		go func() {
			fmt.Scanf("%s\n", &answer)
			aChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case answer := <-aChannel:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("\nYou scored %d from %d", correct, len(problems))
}
