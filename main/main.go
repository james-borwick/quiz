package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/james-borwick/quiz/questions"
)

func main() {
	var bitsQA map[string]string
	var bitsData, _ = ioutil.ReadFile("bits.json")
	json.Unmarshal(bitsData, &bitsQA)

	var binaryQA map[string]string
	var binaryData, _ = ioutil.ReadFile("binary.json")
	json.Unmarshal(binaryData, &binaryQA)

	var choice string
	var correctAnswer = make([]bool, 0)
	var howManyRight int

	fmt.Print(`┏━━━━━━━━━━━━━━┓
┃ Network Quiz ┃
┗━━━━━━━━━━━━━━┛
`)

	for choice != "q" {
		fmt.Print(`[1] Question
[q] Quit
Choose an option: `)
		fmt.Scan(&choice)
		result := questions.Question()
		correctAnswer = append(correctAnswer, result)
	}

	for _, v := range correctAnswer {
		if v == true {
			howManyRight++
		}
	}
	fmt.Printf("Score: %v/%v\n", howManyRight, len(correctAnswer))
}
