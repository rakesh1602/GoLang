package main

import "fmt"

func main() {
	i := 54

	p := &i         //points to i
	fmt.Println(*p) //read i through the pointer
	*p = 333        //set i through the pointer
	fmt.Println("new value of i is :", i)
}
