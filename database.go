package main

import (
	"database/sql"
	"log"
	"os"
)

// create required tables
func createTables(db *sql.DB) {
	// check is todo.db file exists
	if _, err := os.Stat("./todo.db"); err == nil {
		return
	}

	createTodosTableSQL := `CREATE TABLE todos (
		"id" integer PRIMARY KEY AUTOINCREMENT,
		"text" TEXT,
		"status" integer
	);`
	statement, err := db.Prepare(createTodosTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

// insert new todo
func insertTodo(db *sql.DB, text string, status bool) {
	insertTodoSQL := `INSERT INTO todos(text, status) VALUES (?, ?)`
	statement, err := db.Prepare(insertTodoSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(text, status)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// get all todos
func fetchTodos(db *sql.DB) {
	row, err := db.Query("SELECT * FROM todos ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var text string
		var status bool
		row.Scan(&id, &text, &status)
		todos = append(todos, Todo{ID: id, Title: text, Status: status})
	}
}