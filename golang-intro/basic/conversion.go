package main

import "fmt"

func main() {
	// declare a variable and convert it
	var int32Value int32 = 128
	var int64Value int64 = int64(int32Value)
	var int8Value int8 = int8(int64Value)

	//print the conversion
	fmt.Println(int32Value)
	fmt.Println(int64Value)
	//note when converting to int8Value, the value is overflow (since the max value of int32 is 127)
	fmt.Println(int8Value)

}
