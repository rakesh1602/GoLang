package main

import "fmt"

func main() {
	x := 5 //single switch

	switch x {

	case 1, 5:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("wednesday")

	case 4:
		fmt.Println("Thursday")
	}
}
