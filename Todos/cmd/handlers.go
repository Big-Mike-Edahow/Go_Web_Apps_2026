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
    
	todos, err := app.Todos.GetAllTodos()
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}
	indexTemplate, _ := template.ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", todos)
}

func (app *Application) createHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		files := []string{
			"./templates/base.html",
			"./templates/create.html",
		}
		createTemplate, _ := template.ParseFiles(files...)
		createTemplate.ExecuteTemplate(w, "base", nil)
	case "POST":
		item := r.FormValue("item")

		msg := &Message{
			Item: r.PostFormValue("item"),
		}

		if !msg.Validate() {
			files := []string{
				"./templates/base.html",
				"./templates/create.html",
			}
			createTemplate, _ := template.ParseFiles(files...)
			createTemplate.ExecuteTemplate(w, "base", msg)
		} else {
			err := app.Todos.Insert(item)
			if err != nil {
				log.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func (app *Application) updateHandler(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(todoId)
	err := app.Todos.Update(id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	todoId := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(todoId)
	err := app.Todos.Delete(id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) aboutHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/about.html",
	}
	aboutTemplate, _ := template.ParseFiles(files...)
	aboutTemplate.ExecuteTemplate(w, "base", nil)
}
