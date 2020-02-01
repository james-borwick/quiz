package questions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Question .
func Question() bool {
	var userAnswer string

	questionsAndAnswers := getJSONMap()
	question, answer := getQandA(questionsAndAnswers)

	fmt.Println(question)
	fmt.Print("Answer: ")
	fmt.Scan(&userAnswer)

	if userAnswer == answer {
		fmt.Println("Correct!")
		return true
	}

	fmt.Println("Incorrect.")
	return false
}

func getJSONMap() map[string]string {
	var qaMap map[string]string
	var jsonFiles []string

	myFileInfo, _ := ioutil.ReadDir(".")

	for _, v := range myFileInfo {
		if strings.Contains(v.Name(), ".json") {
			jsonFiles = append(jsonFiles, v.Name())
		}
	}

	sliceLength := int64(len(jsonFiles))
	randomMap := jsonFiles[randomInt(sliceLength)]
	qaMapData, _ := ioutil.ReadFile(randomMap)
	json.Unmarshal(qaMapData, &qaMap)

	return qaMap
}

func getQandA(qaMap map[string]string) (randomQuestion string, randomAnswer string) {
	var questionSlice []string

	for k := range qaMap {
		questionSlice = append(questionSlice, k)
	}

	sliceLength := int64(len(questionSlice))
	randomQuestion = questionSlice[randomInt(sliceLength)]
	randomAnswer = qaMap[randomQuestion]
	return randomQuestion, randomAnswer
}

func randomInt(randomRange int64) int64 {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(randomRange)
	return randomInt
}
