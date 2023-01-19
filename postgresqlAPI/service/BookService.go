package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rakesh1602/golang-postgres-book/database"
	"github.com/rakesh1602/golang-postgres-book/model"
	"log"
	"net/http"
)

func CreateBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var book model.Book
	json.NewDecoder(request.Body).Decode(&book)
	db := database.Connection()
	err := db.Create(&book)
	if err == nil {
		log.Fatal("Book should not be empty")
	}
	json.NewEncoder(response).Encode(book.ID)
	log.Println("Books added successfully")
	return
}

func GetBooks(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")

	log.Println("Retrieving the book from database")

	var book []model.Book

	db := database.Connection()
	db.Find(&book)

	/*err := db.Find(&book)
	if err !=nil{
		log.Fatal("No book found")
	}*/
	json.NewEncoder(response).Encode(book)

	return
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	db := database.Connection()
	var book model.Book
	db.First(&book, params["id"])
	json.NewDecoder(request.Body).Decode(&book)
	db.Save(&book)
	json.NewEncoder(response).Encode(book)
	return
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	db := database.Connection()
	param := mux.Vars(request)
	var book model.Book
	db.First(&book, param["id"])
	db.Delete(&book)
	json.NewEncoder(response).Encode("Book deleted successfully")

}
