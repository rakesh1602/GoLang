package main

import (
	"github.com/joho/godotenv"
	"github.com/rakesh1602/golang-postgres-book/controller"
	"github.com/rakesh1602/golang-postgres-book/database"
)

func main() {
	godotenv.Load("app.env")
	database.Connection()
	controller.Routers()

}
