package main

import (
	"fmt"
	"io"
	"net/http"
)

const localhost = "https://lco.dev"

func main() {
	fmt.Println("LCO  web request")

	response, err := http.Get(localhost)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of the response is: %T", response)

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

	response.Body.Close()
}
