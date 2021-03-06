package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var value interface{}
	jsonBody := `{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`
	err := json.Unmarshal([]byte(jsonBody), &value)
	if err != nil {
		fmt.Println("the error of unmarshal:", err)
		return
	}
	fmt.Println("value:", value)
}