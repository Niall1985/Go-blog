package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"go-blog/models"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/create.html", "templates/edit.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAllPosts()
	templates.ExecuteTemplate(w, "index.html", posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		models.CreatePost(title, content)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	templates.ExecuteTemplate(w, "create.html", nil)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	post := models.GetPostByID(id)

	if r.Method == http.MethodPost {
		post.Title = r.FormValue("title")
		post.Content = r.FormValue("content")
		models.UpdatePost(post)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	templates.ExecuteTemplate(w, "edit.html", post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	models.DeletePost(id)
	http.Redirect(w, r, "/", http.StatusFound)
}
