// package models

// import (
// 	"encoding/json"
// 	"os"
// 	"sync"
// )

// type Post struct {
// 	ID       int    `json:"id"`
// 	Title    string `json:"title"`
// 	Username string `json:"username"`
// 	Content  string `json:"content"`
// }

// var (
// 	posts     []Post
// 	idCounter = 1
// 	mut       sync.Mutex
// 	filePath  = "posts.json" // Path to the JSON file
// )

// func LoadPosts() error {
// 	mut.Lock()
// 	defer mut.Unlock()

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return nil // File does not exist; return an empty list
// 		}
// 		return err
// 	}
// 	defer file.Close()

// 	decoder := json.NewDecoder(file)
// 	if err := decoder.Decode(&posts); err != nil {
// 		return err
// 	}
// 	if len(posts) > 0 {
// 		idCounter = posts[len(posts)-1].ID + 1
// 	}
// 	return nil
// }

// func savePosts() error {
// 	mut.Lock()
// 	defer mut.Unlock()

// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	encoder := json.NewEncoder(file)
// 	return encoder.Encode(posts)
// }

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

// func CreatePost(title, username, content string) {
// 	post := Post{ID: idCounter, Title: title, Username: username, Content: content}
// 	posts = append(posts, post)
// 	idCounter++
// 	savePosts()
// }

// func UpdatePost(updatedPost *Post) {
// 	for i, post := range posts {
// 		if post.ID == updatedPost.ID {
// 			posts[i] = *updatedPost
// 			savePosts()
// 			break
// 		}
// 	}
// }

// func DeletePost(id int) {
// 	for i, post := range posts {
// 		if post.ID == id {
// 			posts = append(posts[:i], posts[i+1:]...)
// 			savePosts()
// 			break
// 		}
// 	}
// }

package models

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

var (
	posts     []Post
	idCounter = 1
	mut       sync.Mutex
	filePath  = "posts.ndjson" // Path to the NDJSON file
)

func LoadPosts() error {
	mut.Lock()
	defer mut.Unlock()

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist; return an empty list
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var post Post
		if err := json.Unmarshal(scanner.Bytes(), &post); err != nil {
			return err
		}
		posts = append(posts, post)
		if post.ID >= idCounter {
			idCounter = post.ID + 1
		}
	}

	return scanner.Err()
}

func appendPostToFile(post Post) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(post)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(data) + "\n")
	return err
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

func CreatePost(title, username, content string) {
	post := Post{ID: idCounter, Title: title, Username: username, Content: content}
	posts = append(posts, post)
	idCounter++
	appendPostToFile(post)
}

func UpdatePost(updatedPost *Post) {
	for i, post := range posts {
		if post.ID == updatedPost.ID {
			posts[i] = *updatedPost
			saveAllPosts() // Save all posts since updating a specific post requires rewriting the file
			break
		}
	}
}

func DeletePost(id int) {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			saveAllPosts() // Save all posts since deleting a specific post requires rewriting the file
			break
		}
	}
}

func saveAllPosts() error {
	mut.Lock()
	defer mut.Unlock()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, post := range posts {
		data, err := json.Marshal(post)
		if err != nil {
			return err
		}
		_, err = writer.WriteString(string(data) + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
