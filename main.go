package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	log.Println("Initial Check Passed...Under the main Function...")
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Book1", Author: "Author1", Year: "2020"},
		Book{ID: 2, Title: "Book2", Author: "Author2", Year: "2019"},
		Book{ID: 3, Title: "Book3", Author: "Author3", Year: "2018"},
		Book{ID: 4, Title: "Book4", Author: "Author4", Year: "2017"},
		Book{ID: 5, Title: "Book5", Author: "Author5", Year: "2016"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8006", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books is called...")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get a book is called...")
	params := mux.Vars(r)
	log.Println(params)

	i, _ := strconv.Atoi(params["id"])
	log.Println(reflect.TypeOf(i))

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add books is called...")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	log.Println(json.NewDecoder(r.Body).Decode(&book))
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a book is called...")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)

}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove a book is called...")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)

}
