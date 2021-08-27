package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func connectToDatabase() *sql.DB {
	database, _ := sql.Open("sqlite3", "./todo.db")
	return database
}

// create required tables
func createTables() {
	db := connectToDatabase()
	defer db.Close()
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

func databaseUpdateTodo(ID int, Title string) {
	db := connectToDatabase()
	defer db.Close()
	updateTodoSQL := `UPDATE todos SET text=? WHERE id=?`
	statement, err := db.Prepare(updateTodoSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(Title, ID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func databaseDeleteTodo(ID int) {
	db := connectToDatabase()
	defer db.Close()
	deleteTodoSQL := `DELETE FROM todos WHERE id=?`
	statement, err := db.Prepare(deleteTodoSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(ID)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func emptyTodos() {
	db := connectToDatabase()
	defer db.Close()
	emptyTodosSQL := `DELETE FROM todos`
	statement, err := db.Prepare(emptyTodosSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
	todos = nil
}

// insert new todo
func insertTodo(text string, status bool) {
	db := connectToDatabase()
	defer db.Close()
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
func fetchTodos() {
	db := connectToDatabase()
	defer db.Close()
	row, err := db.Query("SELECT * FROM todos ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	// empty todos slice
	todos = nil

	// fill it with the new data
	for row.Next() {
		var id int
		var text string
		var status bool
		row.Scan(&id, &text, &status)
		todos = append(todos, Todo{ID: id, Title: text, Status: status})
	}
}
