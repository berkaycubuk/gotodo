package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// homepage for the api
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getTodos")
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(todos)
	responseJson, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/todos", getTodos)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
