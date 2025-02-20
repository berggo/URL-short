package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	key, err := ioutil.ReadFile("hello.asc")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(key))
	fmt.Fprint(w, "\n\nHello")
}

func main() {
	http.HandleFunc("/", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
