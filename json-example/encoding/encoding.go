package main

import (
	"encoding/json"
	"fmt"
)

// Member -
type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// Go data
	member := Member{"Alex", 10, true}

	// Json encoding
	jsonBytes, error := json.Marshal(member)
	if error != nil {
		panic(error)
	}

	// Json byte to string
	jsonString := string(jsonBytes)

	fmt.Println(jsonString)
}
