package post

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// JSONRequest -
func JSONRequest() []byte {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Alex", 10}
	personBytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(personBytes)
	response, error := http.Post("http://httpbin.org/post", "application/json", buff)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	// check Response
	responseBody, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	return responseBody
}
