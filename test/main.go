package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var qaMap map[string]string
	var jsonFiles []string

	myFileInfo, _ := ioutil.ReadDir(".")

	for _, v := range myFileInfo {
		if strings.Contains(v.Name(), ".json") {
			jsonFiles = append(jsonFiles, v.Name())
		}
	}

	qaMapData, _ = ioutil.ReadFile(jsonFiles[randomInt])
	json.Unmarshal(bitsData, &qaMap)

	fmt.Println(jsonFiles)

}
