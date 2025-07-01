package main

import "fmt"

func logging() {
	fmt.Println("Selesai melakukan logging")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")

	// simulasi terjadi error: index out of bound -- defer akan tetap dijalankan
	panic("terjadi error") // dibahas berikutnya
}

func main() {
	runApplication()
}
