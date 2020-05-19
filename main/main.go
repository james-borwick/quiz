package main

import (
	"fmt"

	"github.com/james-borwick/quiz/questions"
)

func main() {
	var correctAnswer = make([]bool, 0)
	var howManyRight int

	fmt.Print(`┏━━━━━━━━━━━━━━┓
┃ Network Quiz ┃
┗━━━━━━━━━━━━━━┛
Type 'quit' to quit.
`)

	for numberOfQuestions := 0; numberOfQuestions < 100; numberOfQuestions++ {
		result, quit := questions.Question()
		if quit == true {
			break
		}
		correctAnswer = append(correctAnswer, result)
	}

	for _, v := range correctAnswer {
		if v == true {
			howManyRight++
		}
	}
	fmt.Printf("Score: %v/%v\n", howManyRight, len(correctAnswer))
}
