package json

import (
	"encoding/json"
)

// Decoding -
func Decoding(Name string, Age int, Active bool) string {
	type Member struct {
		Name   string
		Age    int
		Active bool
	}
	jsonBytes, _ := json.Marshal(Member{Name, Age, Active})

	// Json decode
	var member Member
	error := json.Unmarshal(jsonBytes, &member)
	if error != nil {
		panic(error)
	}

	// member struct field access
	// fmt.Println(member.Name, member.Age, member.Active)
	return string(jsonBytes)
}
