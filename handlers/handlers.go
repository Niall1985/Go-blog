package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	// "log"
	"go-blog/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAllPosts()
	templates.ExecuteTemplate(w, "index.html", posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		username := r.FormValue("username")
		content := r.FormValue("content")
		models.CreatePost(title, username, content)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		templates.ExecuteTemplate(w, "create.html", nil)
	}
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		title := r.FormValue("title")
		username := r.FormValue("username")
		content := r.FormValue("content")
		post := &models.Post{
			ID:       id,
			Title:    title,
			Username: username,
			Content:  content,
		}
		models.UpdatePost(post)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		post := models.GetPostByID(id)
		templates.ExecuteTemplate(w, "edit.html", post)
	}
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	models.DeletePost(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
