Sure! Here is a README file for your Go blog project:

---

# Go Blog

Go Blog is a simple blogging platform built with Go (Golang). It allows users to create, read, update, and delete blog posts. This project demonstrates basic web application development with Go, including routing, handling form submissions, and using HTML templates for rendering views.

## Features

- Create new blog posts
- View a list of all blog posts
- Edit existing blog posts
- Delete blog posts

## Prerequisites

- Go (Golang) installed on your machine (version 1.16+ recommended)

## Project Structure

```
go-blog/
│
├── handlers/
│   ├── handlers.go
│   └── handlers_test.go (optional)
│
├── static/
│   └── styles.css
│
├── templates/
│   ├── index.html
│   ├── create.html
│   └── edit.html
│
├── go.mod
├── go.sum
└── main.go
```

## Getting Started

1. **Clone the repository**:

    ```sh
    git clone https://github.com/Niall1985/go-blog.git
    cd go-blog
    ```

2. **Initialize the project**:

    Ensure you are in the project directory and initialize the Go module:

    ```sh
    go mod init go-blog
    go mod tidy
    ```

3. **Run the application**:

    ```sh
    go run main.go
    ```

    You should see a message indicating that the server has started:

    ```
    Server started at http://localhost:8080
    ```

4. **Open your browser** and navigate to `http://localhost:8080` to see the blog application.

## Handlers

The handlers are responsible for processing HTTP requests and rendering the appropriate HTML templates.

- `HomeHandler`: Renders the list of blog posts.
- `CreatePostHandler`: Handles the creation of new blog posts.
- `EditPostHandler`: Handles the editing of existing blog posts.
- `DeletePostHandler`: Handles the deletion of blog posts.

## HTML Templates

HTML templates are used to render the views for the blog application.

- `index.html`: Displays the list of blog posts and links to create, edit, or delete posts.
- `create.html`: Form for creating a new blog post.
- `edit.html`: Form for editing an existing blog post.

## CSS Styling

The `styles.css` file contains the CSS rules to style the application. It ensures that the blog posts are displayed in a clean and attractive layout.

## Example Usage

1. **Create a new post**:
    - Click on the "Create New Post" link.
    - Fill in the form and submit it to create a new blog post.

2. **Edit an existing post**:
    - Click on the "Edit" link next to a post.
    - Modify the post details and submit the form to save changes.

3. **Delete a post**:
    - Click on the "Delete" link next to a post to remove it.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request. Contributions are welcome!

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE. See the `LICENSE` file for details.

---

