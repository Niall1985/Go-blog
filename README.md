# Go Blog

A simple blog platform built with Go, where users can create, read, update, and delete blog posts. This project uses JSON files for data storage.

## Project Structure

```
go-blog/
│
├── main.go
├── models/
│   └── posts.go
├── handlers/
│   └── handlers.go
├── static/
│   ├── styles.css
│
├── templates/
│   ├── index.html
│   ├── create.html
│   └── edit.html
└── posts.json
```

## Features

- **Create Blog Posts**: Add new posts with titles and content.
- **Read Blog Posts**: View a list of all blog posts.
- **Edit Blog Posts**: Modify existing posts.
- **Delete Blog Posts**: Remove posts.

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/Niall1985/Go-blog.git
   cd go-blog
   ```

2. **Install Dependencies**

   Ensure Go is installed on your system. Navigate to the project directory and install any required dependencies.

3. **Run the Application**

   ```bash
   go run main.go
   ```

   The server will start and listen on port `8080` by default. Access the application at `http://localhost:8080`.

## Project Details

### `main.go`

Sets up the HTTP server and routes. Loads and saves blog posts using the `models` package.

### `models/posts.go`

Handles blog posts data:
- **Loading**: Reads blog posts from `posts.json` into memory.
- **Saving**: Writes blog posts to `posts.json`.
- **CRUD Operations**: Functions for creating, reading, updating, and deleting posts.

### `handlers/handlers.go`

Manages HTTP requests and responses:
- **HomeHandler**: Displays all blog posts.
- **CreatePostHandler**: Handles creation of new blog posts.
- **EditPostHandler**: Manages editing of existing blog posts.
- **DeletePostHandler**: Deletes blog posts.

### `templates/`

Contains HTML templates for rendering blog pages:
- **index.html**: Displays the list of blog posts.
- **create.html**: Form for creating new posts.
- **edit.html**: Form for editing existing posts.

### `static/`

Holds static files like CSS for styling.

### `posts.json`

JSON file where blog posts are stored. The application reads from and writes to this file to persist data across sessions.

## Example JSON Structure

Here's an example of how the `posts.json` file might look:

```json
[
  {
    "id": 1,
    "title": "First Blog Post",
    "content": "This is the content of the first blog post."
  },
  {
    "id": 2,
    "title": "Second Blog Post",
    "content": "This is the content of the second blog post."
  }
]
```

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE. See the [LICENSE](LICENSE) file for details.

## Contributions

Feel free to submit pull requests and open issues for any enhancements or bug fixes.

---
