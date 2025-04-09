package handlers

import (
	"book-management/db"
	"book-management/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	db.DB.Create(&book)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	db.DB.Find(&books)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book

	result := db.DB.First(&book, params["id"])
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book

	result := db.DB.First(&book, params["id"])
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&book)
	db.DB.Save(&book)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book

	result := db.DB.First(&book, params["id"])
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	db.DB.Delete(&book)

	w.WriteHeader(http.StatusNoContent)
}
