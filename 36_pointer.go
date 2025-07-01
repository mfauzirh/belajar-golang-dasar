package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Bandung Baru"

	fmt.Println(address1) // tidak berubah karena pass by value
	fmt.Println(address2)

	address3 := Address{"Sukabumi", "Jawa Barat", "Indonesia"}
	address4 := &address3 // pass by reference

	address4.City = "Sukabumi Baru"

	fmt.Println(address3) // data berubah
	fmt.Println(address4)

	address5 := Address{"Sukabumi", "Jawa Barat", "Indonesia"}
	address6 := &address5 // pass by reference

	address6.City = "Sukabumi Baru"

	*address6 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // merubah data yang ditunjuk oleh pointer

	fmt.Println(address5) // data berubah
	fmt.Println(address6)

}
