package main

import "fmt"

func main() {
	day := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	slice1 := day[2:4] // from n to m (n inclusive, m exclusive)
	fmt.Println(slice1)

	slice2 := day[:5]
	fmt.Println(slice2)

	slice3 := day[5:]
	fmt.Println(slice3)

	slice4 := day[:]
	fmt.Println(slice4)

	/* function in slice */
	fmt.Println(len(slice1))

	fmt.Println(cap(slice1))

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	// Karena daySlice point ke array days, jadi elemen days ikut terubah (Konsep pointer)
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru") // Karena capacity sudah bisa menampung, maka akan membuat array baru
	daySlice2[0] = "Sabtu Lama"                  // karena arraynya baru berbeda dengan daySlice1, jadi tidak efek kesana perubahannya
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	// Membuat slice dari awal
	newSlice := make([]string, 2, 5) // type, length, capacity
	newSlice[0] = "Paw"
	newSlice[1] = "Paw"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Paw")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Iki"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Warning
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
