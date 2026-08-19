// handlers.go
// HTTP Route Handlers

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type PageData struct {
	Post     *Post
	Comments []Comment
}

func (app *Application) indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "Go")
    
	posts, err := app.Posts.GetAllPosts()
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/index.html",
	}

	indexTemplate, _ := template.ParseFiles(files...)
	indexTemplate.ExecuteTemplate(w, "base", posts)
}

func (app *Application) saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content")

	err := app.Posts.Insert(title, content)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Application) viewHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	post, err := app.Posts.GetOnePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comments, err := app.Comments.GetComments(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Post:     post,
		Comments: comments,
	}

	files := []string{
		"./templates/base.html",
		"./templates/view.html",
	}

	viewTemplate, _ := template.ParseFiles(files...)
	viewTemplate.ExecuteTemplate(w, "base", data)
}

func (app *Application) editHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	post, err := app.Posts.GetOnePost(id)
	if err != nil {
		log.Println(err)
	}

	files := []string{
		"./templates/base.html",
		"./templates/edit.html",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
	}
	err = tmpl.ExecuteTemplate(w, "base", post)
	if err != nil {
		log.Print(err.Error())
	}
}

func (app *Application) updateHandler(w http.ResponseWriter, r *http.Request) {
	formPostId := r.PostFormValue("id")
	id, err := strconv.Atoi(formPostId)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("content")

	err = app.Posts.Update(id, title, content)
	if err != nil {
		log.Println(err)
		}
	
	http.Redirect(w, r, "/view?id="+formPostId, http.StatusSeeOther)
}

func (app *Application) commentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		formPostId := r.FormValue("post_id")
		postId, err := strconv.Atoi(formPostId)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		content := r.FormValue("content")

		err = app.Comments.InsertComment(postId, content)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/view?id="+formPostId, http.StatusSeeOther)
	}
}

func (app *Application) deleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = app.Posts.Delete(id)
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
