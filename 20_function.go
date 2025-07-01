package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloWithName(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "John", "Doe"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "middle"
	lastName = "Doe"

	return firstName, middleName, lastName
}

func main() {
	sayHello()
	sayHelloWithName("Eko")
	greeting := getHello("Eko")
	fmt.Println(greeting)

	firstname, lastname := getFullName()
	fmt.Println(firstname, lastname)

	_, last := getFullName() // kita bisa menghiraukan return value dengan menggunakan _
	fmt.Println(last)

	firstName, _, lastName := getCompleteName()
	fmt.Println(firstName, lastName)
}
