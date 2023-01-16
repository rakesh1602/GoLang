package main

import "fmt"

func main() {
	var forMapp = map[string]string{"Name": "Rakesh", "Age": "23"}
	fmt.Println(forMapp)

	//Map using make function

	var a = make(map[string]string)
	a["brand"] = "Toyota"
	a["year"] = "1996"

	fmt.Println(a)

}
