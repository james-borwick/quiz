package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/james-borwick/quiz/randomquestion"
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
		fmt.Print(`[1] Binary
[2] Bits
[q] Quit
Choose a category: `)
		fmt.Scan(&choice)
		if choice == "1" {
			result := randomquestion.Question(binaryQA)
			correctAnswer = append(correctAnswer, result)
		} else if choice == "2" {
			result := randomquestion.Question(bitsQA)
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
