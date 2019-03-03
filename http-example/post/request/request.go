package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	// http.Post example
	requestBody := bytes.NewBufferString("Post plain text")
	response, error := http.Post("http://httpbin.org/post", "text/plain", requestBody)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	// check Response
	responseBody, error := ioutil.ReadAll(response.Body)
	if (error) == nil {
		str := string(responseBody)
		println(str)
	}
}
