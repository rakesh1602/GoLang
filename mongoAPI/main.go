package main

import (
	"log"
	"mongoAPI/router"
	"net/http"
)

func main() {
	log.Print("MongoDB API")

	r := router.Router()
	log.Println("Server is getting started.....")
	log.Fatal(http.ListenAndServe(":4000", r))
	log.Println("Listening at port 4000")
}
