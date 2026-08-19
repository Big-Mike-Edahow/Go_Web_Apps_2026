// main.go
// Go - Items

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Id      int
	Title   string
	Body    string
	Created time.Time
}
type DataModel struct {
	DB *sql.DB
}
type Application struct {
	Items interface {
		GetAllItems() ([]Item, error)
		GetOneItem(id int) (*Item, error)
		Insert(title string, body string) error
		Update(id int, title string, body string) error
		Delete(id int) error
	}
}

func main() {
	var err error
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := &Application{
		Items: &DataModel{DB: db},
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/add", app.addHandler)
	http.HandleFunc("/edit", app.editHandler)
	http.HandleFunc("/delete", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)

	log.Printf("Starting HTTP Server on port:8000")
	log.Fatal(http.ListenAndServe(":8000", logRequest(http.DefaultServeMux)))
}
