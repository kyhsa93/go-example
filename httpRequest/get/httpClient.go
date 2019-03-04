package get

import (
	"io/ioutil"
	"net/http"
)

// HTTPClient -
func HTTPClient() string {
	// create request object
	request, error := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	if error != nil {
		panic(error)
	}

	// Add header
	request.Header.Add("User-Agent", "Crawler")

	// run Request from client object
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	// print result
	bytes, _ := ioutil.ReadAll(response.Body)
	str := string(bytes) // byte to string
	// fmt.Println(str)
	return str
}
