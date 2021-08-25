package main

import (
	"fmt"
	"log"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
