package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp() {
	defer endApp()

	panic("ERROR")
}

func main() {
	runApp()
}
