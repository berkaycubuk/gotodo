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
	responseJson, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

type NewTodoRequest struct {
	Title string `json:"title"`
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: addTodo")
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		return
	case "POST":
		var newTodo NewTodoRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&newTodo)
		insertTodo(newTodo.Title, false)
		fetchTodos()
		responseJson, err := json.Marshal(todos)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(responseJson)
	default:
		responseJson, err := json.Marshal(map[string]string{"message": "Please use POST method"})
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
	}
}

type DeleteTodoRequest struct {
	ID int `json:"id"`
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteTodo")
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		return
	case "POST":
		var todoToDelete DeleteTodoRequest
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&todoToDelete)
		databaseDeleteTodo(todoToDelete.ID)
		fetchTodos()
		responseJson, err := json.Marshal(todos)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseJson)
	default:
		responseJson, err := json.Marshal(map[string]string{"message": "Please use POST method"})
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
	}
}

func deleteTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteTodos")
	setupResponse(&w, r)

	switch r.Method {
	case "OPTIONS":
		return
	case "DELETE":
		emptyTodos()
		responseJson, err := json.Marshal(todos)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(responseJson)
	default:
		responseJson, err := json.Marshal(map[string]string{"message": "Please use DELETE method"})
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseJson)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/todos", getTodos)
	http.HandleFunc("/new", addTodo)
	http.HandleFunc("/delete", deleteTodo)
	http.HandleFunc("/delete-all", deleteTodos)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
