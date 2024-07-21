package main

import (
	"log"
	"net/http"

	"go-blog/handlers"
	"go-blog/models"
)

func main() {
	// Load posts from JSON file
	if err := models.LoadPosts(); err != nil {
		log.Fatalf("Error loading posts: %v", err)
	}

	// Define the port you want to use
	port := ":8080"

	// Set up the route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/post", handlers.CreatePostHandler)
	http.HandleFunc("/post/edit", handlers.EditPostHandler)
	http.HandleFunc("/post/delete", handlers.DeletePostHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Log the server start message with the full URL
	log.Printf("Server started at http://localhost%s\n", port)

	// Start the server and log any errors
	log.Fatal(http.ListenAndServe(port, nil))
}
