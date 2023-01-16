package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Hello")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	log.Print("Hello MOD users")
}

func serveHome(w http.ResponseWriter, t *http.Request) {
	w.Write([]byte("<h1>Welcome to web request</h1>"))
}
