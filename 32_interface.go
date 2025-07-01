package main

import "fmt"

// Inteface
type HasName interface {
	GetName() string
}

func SayHello(hashName HasName) {
	fmt.Println("Hello", hashName.GetName())
}

// Person dianggap mengimplementasikan HasName karena sesuai kontrak
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// Begitupula dengan anima
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "John"}
	SayHello(person)

	animal := Animal{Name: "Lion"}
	SayHello(animal)
}
