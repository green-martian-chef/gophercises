package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tscangussu/gophercises/cyoa/parser"
)

// StoryPrinter prints a story to the screen
func StoryPrinter(s parser.Story, c string) {
	chapter := c
	chapters := len(s[chapter].Options)
	answer := ""

	for chapters > 0 {
		ChapterPrinter(s, chapter)

		fmt.Print("Type a number to choose the next chapter: ")
		fmt.Scanf("%s", &answer)

		arc, err := strconv.Atoi(answer)
		if err != nil {
			log.Fatal(err)
		}
		chapter = s[chapter].Options[arc-1].Arc
		chapters = len(s[chapter].Options)
	}

	ChapterPrinter(s, chapter)
	fmt.Println("The End")

}

// ChapterPrinter prints a chapter to the screen
func ChapterPrinter(s parser.Story, c string) {
	fmt.Println()
	fmt.Println(s[c].Title)
	fmt.Println()
	fmt.Println("---")
	fmt.Println()
	for _, v := range s[c].Story {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println("---")
	fmt.Println()
	for i, v := range s[c].Options {
		fmt.Printf("%v: %s\n", i+1, v.Text)
	}
	fmt.Println()
}
