package post

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// FormRequest -
func FormRequest() []byte {
	// http.PostForm example
	response, error := http.PostForm("http://httpbin.org/post", url.Values{"Name": {"Lee"}, "Age": {"10"}})
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
