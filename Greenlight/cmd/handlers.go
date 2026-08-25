// handlers.go
// HTTP Route Handlers

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Info struct {
	Status      string
	Environment string
	Version     string
}

func (app *Application) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	movies, err := app.Movies.GetAllMovies()
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/partials/nav.html",
		"./templates/index.html",
	}
	indexTemplate, err := template.New("index.html").Funcs(app.funcMap).ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", movies)
}

func (app *Application) viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		log.Println(err)
	}

	movie, err := app.Movies.GetOneMovie(id)
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/partials/nav.html",
		"./templates/view.html",
	}
	viewTemplate, _ := template.New("view.html").Funcs(app.funcMap).ParseFiles(files...)
	viewTemplate.ExecuteTemplate(w, "base", movie)
}

func (app *Application) infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	info := Info{
		Status:      app.config.Status,
		Environment: app.config.Env,
		Version:     version,
	}

	files := []string{
		"./templates/base.html",
		"./templates/partials/nav.html",
		"./templates/info.html",
	}
	infoTemplate, _ := template.New("info.html").Funcs(app.funcMap).ParseFiles(files...)
	infoTemplate.ExecuteTemplate(w, "base", info)
}

func (app *Application) createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	switch r.Method {
	case "GET":
		files := []string{
			"./templates/base.html",
			"./templates/partials/nav.html",
			"./templates/create.html",
		}

		createTemplate, _ := template.New("create.html").Funcs(app.funcMap).ParseFiles(files...)
		createTemplate.ExecuteTemplate(w, "base", nil)
	case "POST":
		title := r.FormValue("title")
		year, err := strconv.Atoi(r.FormValue("year"))
		if err != nil {
			log.Println(err)
		}
		runtime, err := strconv.Atoi(r.FormValue("runtime"))
		if err != nil {
			log.Println(err)
		}
		genre := r.FormValue("genre")

		id, err := app.Movies.Insert(title, year, runtime, genre)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *Application) editHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			log.Println(err)
		}

		movie, err := app.Movies.GetOneMovie(id)
		if err != nil {
			log.Println(err)
		}

		files := []string{
			"./templates/base.html",
			"./templates/partials/nav.html",
			"./templates/edit.html",
		}

		editTemplate, _ := template.New("edit.html").Funcs(app.funcMap).ParseFiles(files...)
		editTemplate.ExecuteTemplate(w, "base", movie)
	case "POST":
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println(err)
		}
		title := r.FormValue("title")
		year, err := strconv.Atoi(r.FormValue("year"))
		if err != nil {
			log.Println(err)
		}
		runtime, err := strconv.Atoi(r.FormValue("runtime"))
		if err != nil {
			log.Println(err)
		}
		genre := r.FormValue("genre")

		err = app.Movies.Update(id, title, year, runtime, genre)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
	}
}

func (app *Application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		log.Println(err)
	}

	err = app.Movies.Delete(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (app *Application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./templates/base.html",
		"./templates/partials/nav.html",
		"./templates/about.html",
	}

	aboutTemplate, _ := template.New("about.html").Funcs(app.funcMap).ParseFiles(files...)
	aboutTemplate.ExecuteTemplate(w, "base", nil)
}
