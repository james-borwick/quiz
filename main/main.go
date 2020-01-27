package main

import (
	"encoding/json"
	"fmt"

	"github.com/james-borwick/quiz/binary"
	"github.com/james-borwick/quiz/bits"
	"github.com/james-borwick/quiz/randomquestion"
)

var bitsQA map[string]string



json.Unmarshal()

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
[q] Quit
Choose a category: `)
		fmt.Scan(&choice)
		if choice == "1" {
			result := randomquestion.Question(binary.QuestionsAndAnswers)
			correctAnswer = append(correctAnswer, result)
		} else if choice == "2" {
			result := randomquestion.Question(bits.QuestionsAndAnswers)
			correctAnswer = append(correctAnswer, result)
		}
	}
	for _, v := range correctAnswer {
		if v == true {
			howManyRight++
		}
	}
	fmt.Printf("Score: %v/%v\n", howManyRight, len(correctAnswer))
}
