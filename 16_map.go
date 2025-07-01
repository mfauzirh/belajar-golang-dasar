package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Paw",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// create new map
	book := make(map[string]string)
	book["title"] = "Learning Golang"
	book["author"] = "Muhammad Fauzi"
	book["wrong"] = "Ups"

	fmt.Println(len(book))      // get length of map
	fmt.Println(book["author"]) // access with key

	delete(book, "wrong") // delete key from map
	fmt.Println(book)
}
