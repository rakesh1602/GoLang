package main

import "fmt"

func myFunction() {
	fmt.Println("My function executed")
}

//With parameters

func withParameter(fname string) {
	fmt.Println("Hello ", fname)
}

//with Return Type

func withReturn(x, y int) int {
	return x + y

}

func hello(x int, y string) (result int, text string) {
	result = x + x
	text = "Hello" + y

	return
}

func person(firstName string, lastName string) string {
	return firstName + lastName
}

func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	myFunction()
	withParameter("Rakesh")
	person("rakesh", "chavan")
	fmt.Println(withReturn(2, 2))

	fmt.Println(hello(4, "Rakesh"))

	region, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region, continent)

}
