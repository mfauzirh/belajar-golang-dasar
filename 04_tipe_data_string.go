package main

import "fmt"

func main() {
	fmt.Println("Ini adalah string")

	var nama string = "Budi"
	fmt.Println("Pajang nama: ", len(nama)) // mendapatkan panjang string

	fmt.Println("Karakter pertama pada nama: ", nama[0]) // mendapatkan karakter pada index ke-n
}
