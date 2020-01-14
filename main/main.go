package main

import (
	"fmt"

	"github.com/james-borwick/quiz/one"
	"github.com/james-borwick/quiz/two"
)

func main() {
	var choice string
	var correctAnswer = make([]bool, 0)
	var howManyRight int
	fmt.Print(`┏━━━━━━━━━━━━━━┓
┃ Network Quiz ┃
┗━━━━━━━━━━━━━━┛
`)
	for choice != "q" {
		fmt.Print(`[1] Binary
[2] TCP
[3] IPv4
[q] Quit
Choose a category: `)
		fmt.Scan(&choice)
		if choice == "1" {
			result := one.Question()
			correctAnswer = append(correctAnswer, result)
		} else if choice == "2" {
			two.Question()
		}
	}
	for _, v := range correctAnswer {
		if v == true {
			howManyRight++
		}
	}
	fmt.Printf("Score: %v/%v\n", howManyRight, len(correctAnswer))
}
