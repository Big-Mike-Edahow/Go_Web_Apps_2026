// main.go
// Go - Snippetbox

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"
)

type Snippet struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}
type DataModel struct {
	DB *sql.DB
}
type Application struct {
	Snippets interface {
		GetAllSnippets() ([]Snippet, error)
		GetOneSnippet(id int) (*Snippet, error)
		Insert(title string, content string) error
		Update(id int, title string, content string) error
		Delete(id int) error
	}
}

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := &Application{
		Snippets: &DataModel{DB: db},
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/view/{id}", app.viewHandler)
	http.HandleFunc("/create", app.createHandler)
	http.HandleFunc("/edit/{id}", app.editHandler)
	http.HandleFunc("/delete/{id}", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)

	log.Printf("Starting HTTP Server on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, logRequest(http.DefaultServeMux)))
}
