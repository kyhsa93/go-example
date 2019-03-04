package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":5000", nil)
}
