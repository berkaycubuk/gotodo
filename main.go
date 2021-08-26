package main

import (
	"fmt"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var todos []Todo

// print todos for debugging
func displayTodos() {
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func main() {
	createTables()
	fetchTodos()
	displayTodos()
	handleRequests()
}
