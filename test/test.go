package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(11))
	fmt.Println(rand.Intn(11))
	fmt.Println(rand.Intn(11))
	fmt.Println(rand.Intn(11))
	fmt.Println(rand.Intn(11))
}
