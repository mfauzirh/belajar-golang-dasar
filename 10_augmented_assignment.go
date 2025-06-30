package main

import "fmt"

func main() {
	var a = 10

	a += 10
	fmt.Println(a)

	a -= 10
	fmt.Println(a)

	a *= 10
	fmt.Println(a)

	a /= 10
	fmt.Println(a)

	a %= 10
	fmt.Println(a)
}
