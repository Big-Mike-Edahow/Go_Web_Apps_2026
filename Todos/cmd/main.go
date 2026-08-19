// main.go
// Go - Todos

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

// Defines what a single todo looks like.
type Todo struct {
	Id        int
	Item      string
	Completed bool
	Created   time.Time
}

// Holds the database connection (*sql.DB) to run queries.
type DataModel struct {
	DB *sql.DB
}

// Acts as the main hub for the app. Todos: An interface with rules
// for getting, adding, updating, and deleting tasks.
type Application struct {
	Todos interface {
		GetAllTodos() ([]Todo, error)
		Insert(item string) error
		Update(id int) error
		Delete(id int) error
	}
}

func main() {
	// Initialize the database connection.
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Creates the main app object. It connects the Todos field
	// to the TodoModel database functions.
	app := &Application{
		Todos: DataModel{DB: db},
	}

	// Serve static files.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// HTTP routes.
	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/create", app.createHandler)
	http.HandleFunc("/update", app.updateHandler)
	http.HandleFunc("/delete", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)

	// Start the HTTP server.
	log.Println("Listening and serving on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", logRequest(http.DefaultServeMux)))
}
