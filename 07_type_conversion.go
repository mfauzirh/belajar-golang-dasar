package main

import "fmt"

func main() {
	var int32Num int32 = 32768
	var int64Num int64 = int64(int32Num)

	var int16Num int16 = int16(int32Num) // Warning Overflow

	fmt.Println(int32Num)
	fmt.Println(int64Num)
	fmt.Println(int16Num)

	var name = "John Doe"
	e := name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
