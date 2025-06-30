package main

func main() {
	/* Tipe Data Number: Integer */

	var int8Num int8 = 127 // definisi variable dengan tipe data int8, detail variable akan dibahas nanti
	println(int8Num)

	var int16Num int16 = 32767
	println(int16Num)

	var int32Num int32 = 2147483647
	println(int32Num)

	var int64Num int64 = 9223372036854775807
	println(int64Num)

	var uint8Num = 255
	println(uint8Num)

	var uint16Num uint16 = 65535
	println(uint16Num)

	var uint32Num uint32 = 4294967295
	println(uint32Num)

	var uint64Num uint64 = 18446744073709551615
	println(uint64Num)

	/* Tipe Data Number: Float */

	var float32Num float32 = 3.14
	println(float32Num)

	var float64Num float64 = 3.14
	println(float64Num)

	var complex64Num complex64 = complex64(1.5) // 1.5i
	println(complex64Num)

	var complex128Num complex128 = complex128(1.5) // 1.5i
	println(complex128Num)

	/* Tipe Data Number: Alias */

	var byteNum byte = 255 // alias for uint8
	println(byteNum)

	var runeNum rune = 2147483647 // alias for int32
	println(runeNum)

	var intNum int = 9223372036854775807 // minimal int32
	println(intNum)

	var uintNum uint = 18446744073709551615 // minimal uint32
	println(uintNum)
}
