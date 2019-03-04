package json

import (
	"encoding/json"
)

// Encoding -
func Encoding(Name string, Age int, Active bool) string {
	type Member struct {
		Name   string
		Age    int
		Active bool
	}

	// Go data
	member := Member{Name, Age, Active}

	// Json encoding
	jsonBytes, error := json.Marshal(member)
	if error != nil {
		panic(error)
	}

	// Json byte to string
	jsonString := string(jsonBytes)

	// fmt.Println(jsonString)
	return jsonString
}
