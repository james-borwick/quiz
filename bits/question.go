package bits

import (
	"fmt"
	"math/rand"
	"time"
)

var bitsMap = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
}

var lookup = map[string]string{
	"0": "0",
	"1": "2",
	"2": "4",
	"3": "8",
	"4": "16",
	"5": "32",
	"6": "64",
	"7": "128",
	"8": "256",
}

var stringResult string

// Question for bits.
func Question() bool {
	var answer string
	randomInt := randomSelection()
	fmt.Printf("Question: How many values are possible using %v bits?\n", bitsMap[randomInt])
	fmt.Print("Answer: ")
	fmt.Scan(&answer)
	stringResult = lookup[bitsMap[randomInt]]
	if answer == stringResult {
		fmt.Println("Correct! Well done.")
		return true
	}
	fmt.Println("Incorrect.")
	return false
}

func randomSelection() int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(8)
	return random
}
