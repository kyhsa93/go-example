package post

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// XMLReqeust -
func XMLReqeust() []byte {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{"Alex", 10}
	personBytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(personBytes)
	response, error := http.Post("http://httpbin.org/post", "application/xml", buff)
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
