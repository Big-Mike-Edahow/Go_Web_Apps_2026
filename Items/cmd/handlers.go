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

	items, err := app.Items.GetAllItems()
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}
	indexTemplate, _ := template.ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", items)
}

func (app *Application) addHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		files := []string{
			"./templates/base.html",
			"./templates/add.html",
		}

		addTemplate, _ := template.ParseFiles(files...)
		addTemplate.ExecuteTemplate(w, "base", nil)
	case http.MethodPost:
		title := r.FormValue("title")
		body := r.FormValue("body")

		err := app.Items.Insert(title, body)
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

		item, err := app.Items.GetOneItem(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		files := []string{
			"./templates/base.html",
			"./templates/edit.html",
		}

		editTemplate, _ := template.ParseFiles(files...)
		editTemplate.ExecuteTemplate(w, "base", item)
	case http.MethodPost:
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		body := r.FormValue("body")

		err = app.Items.Update(id, title, body)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *Application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = app.Items.Delete(id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (app *Application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./templates/base.html",
		"./templates/about.html",
	}

	aboutTemplate, _ := template.ParseFiles(files...)
	aboutTemplate.ExecuteTemplate(w, "base", nil)
}
