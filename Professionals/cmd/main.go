// main.go
// Go - Professionals

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Record struct {
	Id      int
	Name    string
	Age     string
	Employ  string
	Created time.Time
}

type DataModel struct {
	DB *sql.DB
}

type Application struct {
	Records interface {
		GetAllRecords() ([]Record, error)
		GetOneRecord(id int) (*Record, error)
		Insert(name string, age string, employ string) error
		Update(id int, name string, age string, employ string) error
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
		Records: &DataModel{DB: db},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/view", app.viewHandler)
	http.HandleFunc("/create", app.createHandler)
	http.HandleFunc("/save", app.saveHandler)
	http.HandleFunc("/edit", app.editHandler)
	http.HandleFunc("/update", app.updateHandler)
	http.HandleFunc("/delete", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)

	log.Println("Listening and serving on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", logRequest(http.DefaultServeMux)))
}
