package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) SayHello() {
	fmt.Println("Hello, my name is", customer.Name)
}

func main() {
	var paw Customer
	paw.Name = "Paw"
	paw.Address = "Bandung"
	paw.Age = 22

	fmt.Println(paw)

	// Membuat struct secara langsung dengan struct literals
	joko := Customer{
		Name:    "Joko",
		Address: "Jogja",
		Age:     30,
	}

	fmt.Println(joko)

	paw.SayHello()
}
