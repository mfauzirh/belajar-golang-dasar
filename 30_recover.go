package main

import "fmt"

func main() {
	var recover = func() {
		message := recover()
		fmt.Println("Terjadi error", message)
	}
	defer recover()
	panic("panic")
}
