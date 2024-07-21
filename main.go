package main

import (
	"go-blog/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/post", handlers.CreatePostHandler)
	http.HandleFunc("/post/edit", handlers.EditPostHandler)
	http.HandleFunc("/post/delete", handlers.DeletePostHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	port := ":8080"

	log.Printf("Server started at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
