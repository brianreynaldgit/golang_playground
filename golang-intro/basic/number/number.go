package main

import (
	"fmt"
)

func main() {
	/// there are 2 kinds of numbers in golang
	//integer int8 int16 int32 int64
	var int8NumMax int8 = 127  // from binary : 0 111 1111
	var int8NumMin int8 = -128 // from binary : -1 000 0000
	fmt.Println(int8NumMin, int8NumMax)

	// the binary representation comes from

	/// below case could not get compiled because of overflows
	//var int8NumMin int8 = -129
	//var int8NumMax int8 = 128

	//there are also unsigned integer uint8 uint16 uint32 uint64
	// var uint8NumMin uint8 = 0 // from binary 0000 0000
	// var uint8NumMax uint8 = 256 // from binary 1111 1111

	// Floating Number : float32 float64
	var float32Num float32 = 123.5
	fmt.Println(float32Num)

	// Complex64 & complex128
	var complex64Num complex64 = 123
	fmt.Println(complex64Num)

	//there are also other types
	var byteNum byte = 123 // this is actually the same as uint8
	var runeNum rune = 123 // this is actually the same as int32
	fmt.Println(byteNum, runeNum)
}
