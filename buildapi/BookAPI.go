package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "go/types"
	"log"
	"net/http"
)

type Authors struct {
	AuthorID       string `json:"authorID"`
	AuthorName     string `json:"authorName"`
	AuthorBookName string `json:"authorBookName"`
}

func main() {
	log.Println("Calling API method")

	bookDetailsByAuthor()

	r := mux.NewRouter()
	r.HandleFunc("/home", authorHome).Methods("GET")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBookById).Methods("GET")

	log.Fatal(http.ListenAndServe(":4301", r))

}

var author []Authors

func bookDetailsByAuthor() {
	author = append(author, Authors{AuthorID: "1", AuthorName: "Alex", AuthorBookName: "The Ross"})
	author = append(author, Authors{AuthorID: "2", AuthorName: "Renny", AuthorBookName: "The IT Guy"})
}

func authorHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the Author Managers</h1>"))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all books of author")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("No books found")
	}
	json.NewEncoder(w).Encode(author)
	return
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting book by id")
	w.Header().Set("content-type", "application/json")

	//grab id from request
	param := mux.Vars(r)

	//loop through the books
	for _, author := range author {
		if author.AuthorID == param["id"] {
			json.NewEncoder(w).Encode(author)
			return
		}
	}
	json.NewEncoder(w).Encode("No book find with id {}")
	return
}
