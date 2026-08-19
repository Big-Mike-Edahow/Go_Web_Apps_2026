// main.go
// Go - Blogger

package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Post struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}

type Comment struct {
	Id      int
	PostId  int
	Content string
	Created time.Time
}

type Application struct {
	Posts interface {
		GetAllPosts() ([]Post, error)
		GetOnePost(id int) (*Post, error)
		Insert(title string, content string) error
		Update(id int, title string, content string) error
		Delete(id int) error
	}
	Comments interface {
		GetComments(id int) ([]Comment, error)
		InsertComment(postId int, content string) error
	}
}

type DataModel struct {
	DB *sql.DB
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := &Application{
		Posts:    &DataModel{DB: db},
		Comments: &DataModel{DB: db},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/save", app.saveHandler)
	http.HandleFunc("/view", app.viewHandler)
	http.HandleFunc("/edit", app.editHandler)
	http.HandleFunc("/update", app.updateHandler)
	http.HandleFunc("/comment", app.commentHandler)
	http.HandleFunc("/delete", app.deleteHandler)
	http.HandleFunc("/about", app.aboutHandler)
	log.Println("Listening and serving HTTP on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", logRequest(http.DefaultServeMux)))
}
