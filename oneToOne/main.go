package main

import (
	"github.com/joho/godotenv"
	"github.com/rakesh1602/GoLang/connection"
	"github.com/rakesh1602/GoLang/controller"
)

func main() {
	godotenv.Load("app.env")
	connection.Connection()
	controller.Routers()
}
