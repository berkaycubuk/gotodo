package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	id     int    `json:"id"`
	title  string `json:"title"`
	status bool   `json:"status"`
}

var Todos []Todo

// homepage for the api
func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// create required tables
func createTables(db *sql.DB) {
	// check is todo.db file exists
	if _, err := os.Stat("./todo.db"); err == nil {
		return
	}

	createTodosTableSQL := `CREATE TABLE todos (
		"id" integer PRIMARY KEY AUTOINCREMENT,
		"text" TEXT
	);`
	statement, err := db.Prepare(createTodosTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

// insert new todo
func insertTodo(db *sql.DB, text string) {
	insertTodoSQL := `INSERT INTO todos(text) VALUES (?)`
	statement, err := db.Prepare(insertTodoSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(text)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func fetchTodos(db *sql.DB) {
	row, err := db.Query("SELECT * FROM todos ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var text string
		row.Scan(&id, &text)
		Todos = append(Todos, Todo{id, text, false})
	}
}

// print todos for debugging
func displayTodos(db *sql.DB) {
	for _, todo := range Todos {
		fmt.Println("Todo: ", todo.title)
	}
}

func main() {
	database, _ := sql.Open("sqlite3", "./todo.db")
	defer database.Close()
	createTables(database)
	fetchTodos(database)

	/*
		insertTodo(database, "Deneme")
		insertTodo(database, "Merhaba")
		insertTodo(database, "Hoooohooo")
	*/

	displayTodos(database)
	// handleRequests()
}
