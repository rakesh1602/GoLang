package controller

import (
	"github.com/gorilla/mux"
	"github.com/rakesh1602/golang-postgres-book/service"
	"net/http"
)

func Routers() {
	router := mux.NewRouter()
	router.HandleFunc("/add-books", service.CreateBook).Methods("POST")
	router.HandleFunc("/get-books", service.GetBooks).Methods("GET")
	router.HandleFunc("/update-book/{id}", service.UpdateBook).Methods("PUT")
	router.HandleFunc("/delete-book/{id}", service.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":9091", router)
}
