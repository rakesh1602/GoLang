package router

import (
	"github.com/gorilla/mux"
	"mongoAPI/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies/{id}", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/addmovie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/update/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", controller.DeleteMovie).Methods("DELETE")

	return router
}
