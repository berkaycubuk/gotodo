package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var todos []Todo

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

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

func handleRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/todos", getTodos)
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
		status := false
		row.Scan(&id, &text)
		todos = append(todos, Todo{ID: id, Title: text, Status: status})
	}
}

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
		insertTodo(database, "Deneme")
		insertTodo(database, "Merhaba")
		insertTodo(database, "Hoooohooo")
	*/
	handleRequests()
}
