/* handlers.go */
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
    
	records, err := app.Records.GetAllRecords()
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}
	indexTemplate, _ := template.ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", records)
}

func (app *Application) viewHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	record, err := app.Records.GetOneRecord(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	files := []string{
		"./templates/base.html",
		"./templates/view.html",
	}

	viewTemplate, _ := template.ParseFiles(files...)
	viewTemplate.ExecuteTemplate(w, "base", record)
}

func (app *Application) createHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/create.html",
	}

	createTemplate, _ := template.ParseFiles(files...)
	createTemplate.ExecuteTemplate(w, "base", nil)
}

func (app *Application) saveHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./templates/base.html",
		"./templates/create.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
	}

	msg := &Message{
		Name:   r.PostFormValue("name"),
		Age:    r.PostFormValue("age"),
		Major: r.PostFormValue("major"),
	}

	if !msg.Validate() {
		err = tmpl.ExecuteTemplate(w, "base", msg)
		if err != nil {
			log.Print(err.Error())
		}
	} else {
		name := r.FormValue("name")
		age := r.FormValue("age")
		major := r.FormValue("major")

		err := app.Records.Insert(name, age, major)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (app *Application) editHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	files := []string{
		"./templates/base.html",
		"./templates/edit.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
	}
	record, err := app.Records.GetOneRecord(id)
	if err != nil {
		log.Println(err)
	}

	msg := &Message{
		Id:		record.Id,
		Name:   record.Name,
		Age:    record.Age,
		Major: record.Major,
	}

	err = tmpl.ExecuteTemplate(w, "base", msg)
	if err != nil {
		log.Print(err.Error())
	}
}

func (app *Application) updateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PostFormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	files := []string{
		"./templates/base.html",
		"./templates/edit.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
	}

	msg := &Message{
		Name:   r.PostFormValue("name"),
		Age:    r.PostFormValue("age"),
		Major: r.PostFormValue("major"),
	}

	if !msg.Validate() {
		err = tmpl.ExecuteTemplate(w, "base", msg)
		if err != nil {
			log.Print(err.Error())
		}
	} else {
		name := r.FormValue("name")
		age := r.FormValue("age")
		major := r.FormValue("major")

		err := app.Records.Update(id, name, age, major)
		if err != nil {
			log.Println(err)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = app.Records.Delete(id)
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
