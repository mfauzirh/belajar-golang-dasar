package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) // panic, karena tidak bisa dilakukan
	// fmt.Println(resultInt)

	// Untuk lebih aman, kita bisa gunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
