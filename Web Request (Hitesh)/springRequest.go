package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func getRequest() {
	const url = "http://localhost:9091/books/3"
	fmt.Println("Getting from spring localhost")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	log.Printf(content)
}

func postRequest() {
	const urlPost = "http://localhost:9091/addbooks"

	requestBody := strings.NewReader(`
			{
  "bookName": "fromgolang",
  "authorList": [
    {
      "authorName": "from golang"
    }
  ]
}
`)

	response, err := http.Post(urlPost, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	log.Printf(string(content))

}

func main() {
	getRequest()
	postRequest()
}
