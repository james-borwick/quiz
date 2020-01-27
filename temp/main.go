package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var bitsQA map[string]string
var bitsData, _ = ioutil.ReadFile("bits.json")

func main() {
	var bitsQA map[string]string
	var bitsData, _ = ioutil.ReadFile("bits.json")
	json.Unmarshal(bitsData, &bitsQA)

	var binaryQA map[string]string
	var binaryData, _ = ioutil.ReadFile("binary.json")
	json.Unmarshal(binaryData, &binaryQA)

	for k, v := range bitsQA {
		binaryQA[k] = v
	}

	fmt.Println(binaryQA)
}
