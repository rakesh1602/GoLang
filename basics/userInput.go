package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	/*	reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter anything")
		myname, _ := reader.ReadString('\n')
		fmt.Println(myname)*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number")
	myRating, _ := reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println(myNumRating)
}
