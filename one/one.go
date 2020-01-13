package one

import (
	"fmt"
	"math/rand"
	"time"
)

var choices = map[int]string{
	1:  "00000001",
	2:  "00000010",
	3:  "00000011",
	4:  "00000100",
	5:  "00000101",
	6:  "00000110",
	7:  "00000111",
	8:  "00010000",
	9:  "00010001",
	10: "00001010",
}

var result = map[string]string{
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

// Question one.
func Question() {
	binary := randomSelection()
	var answer string
	fmt.Printf("Question 1 : What is %s in decimal format?\n", binary)
	fmt.Print("Answer     : ")
	fmt.Scan(&answer)
	if answer == result[binary] {
		fmt.Println("Correct! Well done.")
	} else {
		fmt.Println("Incorrect. Try again.")
	}
}

func randomSelection() string {
	rand.Seed(time.Now().UnixNano())
	return choices[rand.Intn(10)]
}
