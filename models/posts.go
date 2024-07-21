// package models

// type Post struct {
// 	ID      int
// 	Title   string
// 	Content string
// }

// var posts = []Post{}
// var idCounter = 1

// func GetAllPosts() []Post {
// 	return posts
// }

// func GetPostByID(id int) *Post {
// 	for _, post := range posts {
// 		if post.ID == id {
// 			return &post
// 		}
// 	}
// 	return nil
// }

// func CreatePost(title, content string) {
// 	posts = append(posts, Post{ID: idCounter, Title: title, Content: content})
// 	idCounter++
// }

// func UpdatePost(updatedPost *Post) {
// 	for i, post := range posts {
// 		if post.ID == updatedPost.ID {
// 			posts[i] = *updatedPost
// 			break
// 		}
// 	}
// }

// func DeletePost(id int) {
// 	for i, post := range posts {
// 		if post.ID == id {
// 			posts = append(posts[:i], posts[i+1:]...)
// 			break
// 		}
// 	}
// }

package models

import (
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	posts     []Post
	idCounter = 1
	mu        sync.Mutex
	filePath  = "posts.json" // Path to the JSON file
)

func LoadPosts() error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist; return an empty list
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&posts); err != nil {
		return err
	}
	if len(posts) > 0 {
		idCounter = posts[len(posts)-1].ID + 1
	}
	return nil
}

func savePosts() error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(posts)
}

func GetAllPosts() []Post {
	return posts
}

func GetPostByID(id int) *Post {
	for _, post := range posts {
		if post.ID == id {
			return &post
		}
	}
	return nil
}

func CreatePost(title, content string) {
	post := Post{ID: idCounter, Title: title, Content: content}
	posts = append(posts, post)
	idCounter++
	savePosts() // Save posts to JSON file
}

func UpdatePost(updatedPost *Post) {
	for i, post := range posts {
		if post.ID == updatedPost.ID {
			posts[i] = *updatedPost
			savePosts() // Save posts to JSON file
			break
		}
	}
}

func DeletePost(id int) {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			savePosts() // Save posts to JSON file
			break
		}
	}
}
