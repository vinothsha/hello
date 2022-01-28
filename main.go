package main

import (
	"io"
	"net/http"
)

func hello(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Hello vinoth! FINALLY we hit the landing page")
}
func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
