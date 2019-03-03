package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Person -
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Alex", 10}
	personBytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(personBytes)

	// create Request object
	request, error := http.NewRequest("POST", "http://httpbin.org/post", buff)
	if error != nil {
		panic(error)
	}

	// Add Content-type haeder
	request.Header.Add("Content-Type", "application/xml")

	// excute Request in Client
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	// check Response
	responseBody, error := ioutil.ReadAll(response.Body)
	if error == nil {
		str := string(responseBody)
		println(str)
	}
}
