package main

import "fmt"

type Person struct {
	name        string
	age         int
	isAvailable bool
}

func main() {
	var person Person
	person.name = "rakesh"
	person.age = 23
	person.isAvailable = false

	fmt.Println("Name: ", person.name, "Age: ", person.age, "isAvailable: ", person.isAvailable)
	fmt.Println(person)
}
