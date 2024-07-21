package main

import (
	"log"
	"net/http"

	"go-blog/handlers"
	"go-blog/models"
)

func main() {
	if err := models.LoadPosts(); err != nil {
		log.Fatalf("Error loading posts: %v", err)
	}

	port := ":8080"

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/post", handlers.CreatePostHandler)
	http.HandleFunc("/post/edit", handlers.EditPostHandler)
	http.HandleFunc("/post/delete", handlers.DeletePostHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Server started at http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
