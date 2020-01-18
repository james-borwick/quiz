package test

import (
	"fmt"
	"time"
)

// Question time.
func Question() {
	var x time.Time = time.Now()
	fmt.Println(x.Minute())
}
