package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
