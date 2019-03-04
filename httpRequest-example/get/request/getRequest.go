package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// GET request
	response, error := http.Get("http://csharp.news")
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	// print result
	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s\n", string(data))
}
