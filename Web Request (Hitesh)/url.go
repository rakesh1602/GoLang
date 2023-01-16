package main

import (
	"fmt"
	url2 "net/url"
)

const myUrl = "http://localhost:8080/books/1"

func main() {
	fmt.Println(myUrl)

	result, _ := url2.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Print(result.Port())
	fmt.Println(result.RawQuery)

	QParamas := result.Query()

	fmt.Println("The types of the query params are: %T\n", QParamas)
}
