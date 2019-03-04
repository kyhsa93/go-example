package example_test

import (
	"fmt"
	"go-example/httpRequest/get"
	"go-example/httpRequest/post"
	"go-example/json"
	"testing"
)

func TestJsonDecoding(test *testing.T) {
	var Name = "testName"
	var Age = 1
	var Active = true
	data := json.Decoding(Name, Age, Active)
	if data != "{\"Name\":\"testName\",\"Age\":1,\"Active\":true}" {
		test.Error("fale to decode json")
	}
	defer fmt.Println("PASS TestJsonDecoding")
}

func TestJsonEncoding(test *testing.T) {
	var Name = "testName"
	var Age = 1
	var Active = true
	data := json.Encoding(Name, Age, Active)
	if data != "{\"Name\":\"testName\",\"Age\":1,\"Active\":true}" {
		test.Error("fale to encode json")
	}
	defer fmt.Println("PASS TestJsonEncoding")
}

func TestGetRequest(test *testing.T) {
	get.Request()
	defer fmt.Println("PASS TestGetRequest")
}

func TestGetHTTPClient(test *testing.T) {
	get.HTTPClient()
	defer fmt.Println("PASS TestGetHTTPClient")
}

func TestPostFormRequest(test *testing.T) {
	post.FormRequest()
	defer fmt.Println("PASS TestPostFormRequest")
}

func TestPostHTTPClient(test *testing.T) {
	post.HTTPClient()
	defer fmt.Println("PASS TestPostHTTPClient")
}

func TestPostJSONRequest(test *testing.T) {
	post.JSONRequest()
	defer fmt.Println("PASS TestPostJSONRequest")
}

func TestPostRequset(test *testing.T) {
	post.Request()
	defer fmt.Println("PASS TestPostRequset")
}

func TestPostXMLRequest(test *testing.T) {
	post.XMLReqeust()
	defer fmt.Println("PASS TestPostXMLRequest")
}
