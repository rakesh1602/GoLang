package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Website  string   `json:"website"`
	Password string   `json:"_"`
	Tags     []string `json:"tags ,omitempty"`
}

func main() {

	fmt.Print("How to create JSON data")
	encodeJsonData()
}

func encodeJsonData() {
	lcoCourses := []course{
		{"ReactsJs", 299, "rakesh@gmail.com", "Abc123", []string{"web", "go"}},
		{"Angular", 299, "rakesh@gmail.com", "Abc123", []string{"web", "go"}},
		{"spring", 544, "rakesh@gmail.com", "Abc123", nil},
	}

	//Package this data as JSON format

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
