package main

import (
	"fmt"
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, it works!")
}

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":8080", nil)
}
