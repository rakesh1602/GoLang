package main

import "fmt"

func main() {
	//With the fixed length

	var array = [5]int{1, 2, 4, 5, 6}
	fmt.Println(array)

	var array1 = [3]string{"r", "A", "k"}
	fmt.Println(array1)

	//Without the fixed length

	var newArray = [...]int{2, 3, 5, 3, 1, 3, 5, 3}
	fmt.Println(newArray)

	//Element of specific position
	fmt.Println("element if 3rd position ", newArray[3])

	//Change the values of specific element
	newArray[3] = 44
	fmt.Println("after changing the values of specific index", newArray)

	//Array Initialization
	var arr1 = [5]int{}     //not initialized
	var arr2 = [5]int{1, 2} //half initialized

	fmt.Println(arr1)
	fmt.Println(arr2)

	//Length of the Array
	fmt.Println("Length of the new array is ", len(newArray))

}
