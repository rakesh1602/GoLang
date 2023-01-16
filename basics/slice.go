package main

import "fmt"

func main() {
	newslice := []int{1, 2, 4, 4, 2, 4, 2, 4, 2}
	fmt.Println(newslice)
	fmt.Println(cap(newslice))
	fmt.Println(len(newslice))

}
