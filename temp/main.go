package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myInt int64 = 10
	myString := strconv.FormatInt(myInt, 10)
	fmt.Println("myInt:", myInt)
	fmt.Println("myString:", myString)
}