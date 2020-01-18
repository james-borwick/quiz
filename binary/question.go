package binary

import (
	"fmt"
	"math/rand"
	"time"
)

var choices = map[int]string{
	0:  "00000000",
	1:  "00000001",
	2:  "00000010",
	3:  "00000011",
	4:  "00000100",
	5:  "00000101",
	6:  "00000110",
	7:  "00000111",
	8:  "00001000",
	9:  "00001001",
	10: "00001010",
}

var lookup = map[string]string{
	"00000000": "0",
	"00000001": "1",
	"00000010": "2",
	"00000011": "3",
	"00000100": "4",
	"00000101": "5",
	"00000110": "6",
	"00000111": "7",
	"00001000": "8",
	"00001001": "9",
	"00001010": "10",
}

var stringResult string

// Question for binary.
func Question() bool {
	randomInt := randomSelection()
	var answer string
	var valid = false
	for !valid {
		fmt.Printf("Question: What is %s in decimal format?\n", choices[randomInt])
		stringResult = lookup[choices[randomInt]]
		fmt.Print("Answer: ")
		fmt.Scan(&answer)
		valid = validAnswer(answer)
		if !valid {
			fmt.Println("Answer must be within range 1-10. Try again.")
		}
	}
	if answer == stringResult {
		fmt.Println("Correct! Well done.")
		return true
	}
	fmt.Println("Incorrect.")
	return false
}

func randomSelection() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(10)
	return random
}

func validAnswer(s string) bool {
	for _, v := range lookup {
		if s == v {
			return true
		}
	}
	return false
}
