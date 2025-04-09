# 📚 Book Management CRUD API

This is a RESTful API built in **Go** using the **Gorilla Mux** router and a **MySQL** database. It enables you to perform **Create**, **Read**, **Update**, and **Delete (CRUD)** operations on a collection of books.

---

## 🛠️ Technologies Used

- **Golang** - Backend API
- **Gorilla Mux** - HTTP request router
- **MySQL** - Relational database
- **Docker** - Containerization

---

## 🚀 Getting Started

### 📦 Prerequisites

- Go (v1.18+)
- Docker
- Git

### 📁 Clone the Repository

```bash
git clone https://github.com/your-username/book-management-crud-api.git
cd book-management-crud-api
```

> The API will be running at `http://localhost:8080`

---

## 📂 Project Structure

```
.
├── main.go               # Entry point
├── handlers/             # HTTP handlers
├── models/               # Book model & DB logic
├── db/                   # MySQL connection setup
├── docker-compose.yml    # Docker configuration
├── go.mod / go.sum       # Go module files
└── README.md             # Project documentation
```

---

## 📚 API Documentation

### 📘 Book Object

```json
{
  "id": 1,
  "title": "The Pragmatic Programmer",
  "author": "Andrew Hunt",
  "publisher": "Addison-Wesley",
  "year": 1999
}
```

---

## 🔗 Endpoints

### ✅ Get All Books

- **URL**: `/books`
- **Method**: `GET`
- **Response**:

```json
[
  {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "publisher": "Prentice Hall",
    "year": 2008
  }
]
```

---

### 🔍 Get Book by ID

- **URL**: `/books/{id}`
- **Method**: `GET`
- **Response**:

```json
{
  "id": 1,
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "publisher": "Prentice Hall",
  "year": 2008
}
```

---

### ➕ Create a Book

- **URL**: `/books`
- **Method**: `POST`
- **Request Body**:

```json
{
  "title": "Domain-Driven Design",
  "author": "Eric Evans",
  "publisher": "Addison-Wesley",
  "year": 2003
}
```

- **Response**:

```json
{
  "message": "Book created successfully"
}
```

---

### 📝 Update a Book

- **URL**: `/books/{id}`
- **Method**: `PUT`
- **Request Body**:

```json
{
  "title": "Updated Title",
  "author": "Updated Author",
  "publisher": "Updated Publisher",
  "year": 2022
}
```

- **Response**:

```json
{
  "message": "Book updated successfully"
}
```

---

### ❌ Delete a Book

- **URL**: `/books/{id}`
- **Method**: `DELETE`
- **Response**:

```json
{
  "message": "Book deleted successfully"
}
```

---

## 🧪 Testing the API

You can test the endpoints using:

- [Postman](https://www.postman.com/)
- [curl](https://curl.se/)

Example with `curl`:

```bash
curl http://localhost:8080/books
```

---

## 🛢️ Database Schema

```sql
CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255),
    publisher VARCHAR(255),
    year INT
);
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---






