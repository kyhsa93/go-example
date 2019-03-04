package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	http.Handle("/", new(staticHandler))
	http.ListenAndServe(":5000", nil)
}

type staticHandler struct {
	http.Handler
}

func (handler *staticHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// if want use this file path, run command in this directory.
	localPath := "./" + request.URL.Path
	content, error := ioutil.ReadFile(localPath)
	if error != nil {
		response.WriteHeader(404)
		response.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	response.Header().Add("Content-Type", contentType)
	response.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}
	return contentType
}
