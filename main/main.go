package main

import (
	"fmt"
	"github.com/james-borwick/quiz/one"
	"github.com/james-borwick/quiz/two"
)

func main() {
	var choice string
	fmt.Print(`┏━━━━━━━━━━━━━━┓
┃ Network Quiz ┃
┗━━━━━━━━━━━━━━┛
[1] Binary
[2] TCP
[3] IPv4
[4] Quit
Choose a category: `)
	fmt.Scan(&choice)
	switch {
	case choice == "1":
		one.Question()
	case choice == "2":
		two.Question()
	}
}
