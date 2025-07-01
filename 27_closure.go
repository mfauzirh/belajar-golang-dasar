package main

import "fmt"

func main() {
	counter := 0

	// counter dapat diakses dari function, meski ada di luar functionya
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
