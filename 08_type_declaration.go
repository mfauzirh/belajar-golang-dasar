package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJohn NoKTP = "3273441312340091"
	fmt.Println(ktpJohn)
	fmt.Println(NoKTP("1234123412341234"))
}
