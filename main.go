package main

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

// Book struct (Model) (class)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a "slice" (undetermined length arr) Book struct
var books []Book

//what does the asterisk mean?

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)	
}

//Get One Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get Params
	
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

//Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	//Init router
	r := mux.NewRouter();

	//Mock data - @todo- implement DB
	books = append(books, Book{ID: "1", Isbn: "143837662", Title: "Book One", Author: &Author {Firstname: "Dave", Lastname: "Funk"}})
	books = append(books, Book{ID: "2", Isbn: "66666", Title: "Book Two", Author: &Author {Firstname: "Your", Lastname: "Mom"}})

	//Route handlers/Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3001", r))
}
