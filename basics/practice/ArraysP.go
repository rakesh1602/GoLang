package main

import "log"

func main() {
	var newArray = [5]int{1, 2, 4, 5, 5}

	for i := 0; i <= len(newArray); i++ {
		log.Println("print from for loop", newArray)
	}
}
