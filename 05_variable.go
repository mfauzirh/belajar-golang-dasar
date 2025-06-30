package main

func main() {
	var name string // mendefinisikan variable dengan tipe data string

	name = "Paw" // assign value kepada variable

	println(name)

	var age = 12 // tidak perlu mengassign variable, go otomatis menentukan dari nilai disebelah kanan
	println(age)

	isSuccess := true // kita juga bisa tidak menggunakan keyword var, tapi := dengan syarat nilai harus langsung di assign
	println(isSuccess)

	// Di golang juga bisa deklarasi variable secara bersamaan
	var (
		firstName = "Muhammad"
		lastName  = "Fauzi"
	)

	println(firstName + " " + lastName)
}
