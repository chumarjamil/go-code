package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Sample data model
type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books []Book // In memory data storage

func main() {
	// Initialize the router
	router := mux.NewRouter()
	
	// Define routes
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    router.HandleFunc("/api/books", createBook).Methods("POST")
    router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running on PORT: 9090")

	log.Fatal(http.ListenAndServe(":9090", router))
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range books {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            json.NewEncoder(w).Encode(books)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

