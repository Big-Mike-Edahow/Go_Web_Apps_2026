// main.go
// Go - Greenlight

package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Movie struct {
	Id      int
	Title   string
	Year    int
	Runtime int
	Genre   string
	Created time.Time
}

type DataModel struct {
	DB *sql.DB
}

type Config struct {
	Status string
	Port   int
	Env    string
}

type Application struct {
	config  Config
	funcMap template.FuncMap
	logger  *slog.Logger
	Movies  interface {
		GetAllMovies() ([]Movie, error)
		GetOneMovie(id int) (Movie, error)
		Insert(title string, year int, runtime int, genre string) (int, error)
		Update(id int, title string, year int, runtime int, genre string) error
		Delete(id int) error
	}
}

const version = "1.0.0"

func main() {
	var cfg Config
	funcs := template.FuncMap{
		"currentYear": func() int {
			return time.Now().Year()
		},
	}

	flag.IntVar(&cfg.Port, "port", 8000, "API server port")
	flag.StringVar(&cfg.Status, "status", "Avaiable", "Status (Available|Standby|Offline)")
	flag.StringVar(&cfg.Env, "env", "Development", "Environment (Development|Staging|Production)")
	flag.Parse()

	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		config:  cfg,
		funcMap: funcs,
		logger:  logger,
		Movies:  &DataModel{DB: db},
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/{$}", app.indexHandler)
	mux.HandleFunc("/view/{id}", app.viewHandler)
	mux.HandleFunc("/info", app.infoHandler)
	mux.HandleFunc("/create", app.createHandler)
	mux.HandleFunc("/edit/{id}", app.editHandler)
	mux.HandleFunc("/delete/{id}", app.deleteHandler)
	mux.HandleFunc("/about", app.aboutHandler)

	logger.Info("starting server", "addr", cfg.Port, "env", cfg.Env)
	err = http.ListenAndServe(":8000", logRequest(mux))
	logger.Error(err.Error())
	os.Exit(1)
}

