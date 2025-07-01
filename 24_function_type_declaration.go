package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter2(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter2("Eko", spamFilter2)

	filter := spamFilter2
	sayHelloWithFilter2("Anjing", filter)
}
