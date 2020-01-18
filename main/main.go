package main

import (
	"fmt"

	"github.com/james-borwick/quiz/binary"
	"github.com/james-borwick/quiz/bits"
	"github.com/james-borwick/quiz/test"
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
[2] Bits
[3] Time
[q] Quit
Choose a category: `)
		fmt.Scan(&choice)
		if choice == "1" {
			result := binary.Question()
			correctAnswer = append(correctAnswer, result)
		} else if choice == "2" {
			result := bits.Question()
			correctAnswer = append(correctAnswer, result)
		} else if choice == "3" {
			test.Question()
		}
	}
	for _, v := range correctAnswer {
		if v == true {
			howManyRight++
		}
	}
	fmt.Printf("Score: %v/%v\n", howManyRight, len(correctAnswer))
}
