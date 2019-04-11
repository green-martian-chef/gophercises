package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type problem struct {
	q, a string
}

// msgErr takes a string and error, print both messages in the screen and exit.
func msgErr(msg string, err error) {
	fmt.Println(msg, "\n", err)
	os.Exit(1)
}

// parseLines take a slice of []string and returns another slice of []problem.
// Here we put the slice returned by ReadAll() into a struct. Why? If instead
// of a CSV file we had a JSON, YAML or something, we only need to change the
// parser code, the rest remains the same.
func parseLines(lines [][]string) (ret []problem) {
	ret = make([]problem, len(lines))
	for i, p := range lines {
		ret[i] = problem{
			q: p[0],
			a: p[1],
		}
	}
	return
}

// shuffleQuizz returns a shuffled []int if shuffle is true, otherwise returns
// an ordered []int
func shuffleQuiz(s bool, l int, p []problem) []int {
	ret := make([]int, l)

	if s == true {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		ret = r.Perm(l)
	} else {
		for i := range p {
			ret[i] = i
		}
	}
	return ret
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "a CSV file to be parsed in a format 'question,answer'")
	timeLimit := flag.Int("limit", 10, "time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle the quiz order")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		msgErr("The CSV file could not be opened.", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.TrimLeadingSpace = true

	lines, err := r.ReadAll()
	if err != nil {
		msgErr("The CSV file could not be parsed.", err)
	}

	problems := parseLines(lines)
	numProblems := shuffleQuiz(*shuffle, len(problems), problems)

	// NewTimer() creates a new type Timer that will send the current time on
	// its channel after the duration
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// The problem here is you want to ask the user a question while the timer
	// is running but you don't want to accept any answer after the timer has
	// expired. Since NewTimer() returns the time through a channel, you can use
	// select and break the loop when the timer is expired. Select waits on
	// multiple channel operations so you have to put the code that receives the
	// user answer on a goroutine, so if the timer is expired select breaks the
	// loop instead of accept the answer.
	correct := 0
loop:
	for i, v := range numProblems {
		ansChan := make(chan string)
		p := problems[v]

		fmt.Printf("Problem %2d : %s = ", i+1, p.q)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break loop
		case answer := <-ansChan:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("\nYou scored %d from %d", correct, len(problems))
}
