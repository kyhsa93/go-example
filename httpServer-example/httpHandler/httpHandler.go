package main

import (
	"net/http"
)

func main() {
	http.Handle("/", new(testHandler))
	http.ListenAndServe(":5000", nil)
}

// testHandler -
type testHandler struct {
	http.Handler
}

func (handler *testHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	str := "Your Request Path is " + request.URL.Path
	response.Write([]byte(str))
}
