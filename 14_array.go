package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Middle"
	names[2] = "Doe"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat array secara langsung
	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)

	var values2 = [4]int{90, 58, 22, 11}
	fmt.Println(values2)

	fmt.Println("Panjang array ", len(values)) // mendapatkan panjang array
	fmt.Println(values[1])                     // mengakses elemen array menggunakan indeks
}
