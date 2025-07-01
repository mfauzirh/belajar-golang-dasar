package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// Deklrasi anonymous function di variable
	Blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("eko", Blacklist)

	// Menggunakan anonymous function langsung di argument
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
