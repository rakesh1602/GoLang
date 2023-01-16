package main

import "fmt"

// outside the function
var a1 int = 2
var lang string = " go"
var c1 = 333

// By declaring variables
func main() {
	var firstName string = "rakesh"
	var lastName = "chavan"
	x := 34

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(x)

	// Without initializing

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//Assign value after declaration

	var name string
	name = "ramesh"
	fmt.Println("after initializing " + name)

	//Outside the function
	fmt.Println("Outside the function")
	fmt.Println(a1)
	fmt.Println(lang)
	fmt.Println(c1)

	var batman string = "i am batman"
	fmt.Printf("%v ,%T\n", batman, batman)

	//Multiple variables in single line
	var (
		names, age, address = "rakesh", 23, "Vikhroli"
	)

	newName := "ramesh"

	fmt.Println("Multiple variables in sigle line ", names, age, address, newName)

	// variables can be write as a function also

	action := func() {
		fmt.Println("variables as a functions")
	}

	action()

}
