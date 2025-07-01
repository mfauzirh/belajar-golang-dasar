package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}
