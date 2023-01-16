package main

import "fmt"

func main() {
	/*for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("+++---+++")

	for j := 0; j <= 100; j += 10 {
		fmt.Println(j)
	}*/

	//nested loop

	/*	numbers := []int{1, 3, 4, 5, 6}
		var fruits = [3]string{"banana", "apple", "mango"}

		for i := 0; i <= len(numbers); i++ {
			for j := 0; j <= len(fruits); j++ {
				fmt.Println(numbers[i])
				fmt.Println(fruits[j])

			}
		}*/

	//Print both index and values

	var cars = [3]string{"BMW", "MERCEDES", "MARUTI"}
	for idx, val := range cars {
		fmt.Println(idx, val)
	}
}
