// handlers.go
// HTTP Route Handlers

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *Application) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	books, err := app.Books.GetAllBooks()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}

	indexTemplate, _ := template.ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", books)
}

func (app *Application) viewHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	book, err := app.Books.GetOneBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./templates/base.html",
		"./templates/view.html",
	}

	viewTemplate, _ := template.ParseFiles(files...)
	viewTemplate.ExecuteTemplate(w, "base", book)
}

func (app *Application) addHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		files := []string{
			"./templates/base.html",
			"./templates/add.html",
		}

		createTemplate, _ := template.ParseFiles(files...)
		createTemplate.ExecuteTemplate(w, "base", nil)
	case http.MethodPost:
		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		formPrice := r.FormValue("price")
		parsedPrice, err := strconv.ParseFloat(formPrice, 32)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}
		price := float32(parsedPrice)
		excerpt := r.FormValue("excerpt")

		err = app.Books.Insert(isbn, title, author, price, excerpt)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *Application) editHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		book, err := app.Books.GetOneBook(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		files := []string{
			"./templates/base.html",
			"./templates/edit.html",
		}

		editTemplate, _ := template.ParseFiles(files...)
		editTemplate.ExecuteTemplate(w, "base", book)
	case http.MethodPost:
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		formPrice := r.FormValue("price")
		parsedPrice, err := strconv.ParseFloat(formPrice, 32)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}
		price := float32(parsedPrice)
		excerpt := r.FormValue("excerpt")

		err = app.Books.Update(id, isbn, title, author, price, excerpt)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *Application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = app.Books.Delete(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (app *Application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/about.html",
	}

	aboutTemplate, _ := template.ParseFiles(files...)
	aboutTemplate.ExecuteTemplate(w, "base", nil)
}
