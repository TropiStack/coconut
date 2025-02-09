package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
  w.Header().Add("Content-Type", "application/json")
  fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

func main() {
  http.HandleFunc("/", helloWorldHandler)
  http.ListenAndServe(":8080", nil)
}
