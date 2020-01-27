package randomquestion

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomSelection .
func RandomSelection(answersMap map[string]string) (randomQuestion string, randomAnswer string) {
	var questionSlice []string
	
	for k := range answersMap {
		questionSlice = append(questionSlice, k)
	}

	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(int64(len(questionSlice)))
	randomQuestion = questionSlice[randomInt]
	randomAnswer = answersMap[randomQuestion]
	return randomQuestion, randomAnswer
}

// Question .
func Question(QuestionsAndAnswers map[string]string) bool {
	question, answer := RandomSelection(QuestionsAndAnswers)
	var userAnswer string
	fmt.Println(question)
	fmt.Print("Answer: ")
	fmt.Scan(&userAnswer)
	if userAnswer == answer {
		fmt.Println("Correct! Well done.")
		return true
	}
	fmt.Println("Incorrect.")
	return false
}
