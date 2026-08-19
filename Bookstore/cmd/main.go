// main.go
// Go - Bookstore

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Book struct {
	Id      int
	Isbn    string
	Title   string
	Author  string
	Price   float32
	Excerpt string
	Created time.Time
}

type DataModel struct {
	DB *sql.DB
}

type Application struct {
	Books interface {
		GetAllBooks() ([]Book, error)
		GetOneBook(id int) (*Book, error)
		Insert(isbn string, title string, author string, price float32, excerpt string) error
		Update(id int, isbn string, title string, author string, price float32, excerpt string) error
		Delete(id int) error
	}
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := &Application{
		Books: &DataModel{DB: db},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/view", app.viewHandler)
	http.HandleFunc("/add", app.addHandler)
	http.HandleFunc("/edit", app.editHandler)
	http.HandleFunc("/delete", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)

	// Start the HTTP server.
	log.Println("Listening and serving on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", logRequest(http.DefaultServeMux)))
}
