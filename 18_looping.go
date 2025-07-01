package main

import "fmt"

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	names := []string{"Muhammad", "Fauzi", "Rizki"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	book := map[string]string{
		"author": "paw",
		"title":  "another world",
	}

	for key, value := range book {
		fmt.Println("key", key, "=", value)
	}
}
