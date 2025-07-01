package main

import "fmt"

func main() {
	name := "paw"

	if name == "paw" {
		fmt.Println("Hello paw")
	} else if name == "john" {
		fmt.Println("Hello john")
	} else {
		fmt.Println("Hello guest")
	}

	// If dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	grade := "C"

	switch grade {
	case "A":
		fmt.Println("Lulus terbaik")
	case "B":
	case "C":
		fmt.Println("Lulus")
	case "D":
		fmt.Println("Lulus kurang memuaskan")
	default:
		fmt.Println("Tidak lulus")
	}

	// Swith dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
