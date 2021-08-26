package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var todos []Todo

// print todos for debugging
func displayTodos(db *sql.DB) {
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func main() {
	database, _ := sql.Open("sqlite3", "./todo.db")
	defer database.Close()
	createTables(database)
	fetchTodos(database)
	displayTodos(database)

	/*
		insertTodo(database, "Deneme", false)
		insertTodo(database, "Merhaba", false)
		insertTodo(database, "Hoooohooo", true)
	*/
	handleRequests()
}
