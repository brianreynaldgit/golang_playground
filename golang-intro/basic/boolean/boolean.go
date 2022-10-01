package main

import "fmt"

func main() {
	//there are only 2 values for boolean which is true or false
	var boolTrue bool = true
	var boolFalse bool = false
	//the complement of true is false and it works backwards
	if boolTrue == !boolFalse {
		fmt.Println("similar value")
	}
}
