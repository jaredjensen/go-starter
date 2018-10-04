package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book is a book
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author is an author
type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Init mock book data
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, b := range books {
		if b.ID == params["id"] {
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	http.NotFound(w, r)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// Assign an ID and add to collection
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, b := range books {
		if b.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&books[i])
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	http.NotFound(w, r)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, b := range books {
		if b.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "ABCD", Title: "Book 1", Author: &Author{FirstName: "John", LastName: "Smith"}})
	books = append(books, Book{ID: "2", Isbn: "BCDE", Title: "Book 2", Author: &Author{FirstName: "Jane", LastName: "Doe"}})
	books = append(books, Book{ID: "3", Isbn: "CDEF", Title: "Book 3", Author: &Author{FirstName: "Bob", LastName: "Williams"}})

	// Define routes
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Listening at http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
