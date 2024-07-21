package models

type Post struct {
	ID      int
	Title   string
	Content string
}

var posts = []Post{}
var idCounter = 1

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
	posts = append(posts, Post{ID: idCounter, Title: title, Content: content})
	idCounter++
}

func UpdatePost(updatedPost *Post) {
	for i, post := range posts {
		if post.ID == updatedPost.ID {
			posts[i] = *updatedPost
			break
		}
	}
}

func DeletePost(id int) {
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			break
		}
	}
}
