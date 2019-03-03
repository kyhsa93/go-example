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
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	// Json decode
	var member Member
	error := json.Unmarshal(jsonBytes, &member)
	if error != nil {
		panic(error)
	}

	// member struct field access
	fmt.Println(member.Name, member.Age, member.Active)
}
