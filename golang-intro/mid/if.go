package main

import "fmt"

func main() {
	//short statement enables user to write statement first before executing it.
	if age := 5; age < 12 {
		fmt.Println("you are underage")
	} else {
		fmt.Println("hahihuheho")
	}
	fmt.Println("test")
	fmt.Println("test2")
	fmt.Println("master")
	//however, the variable is not accessible outside of the functions.
	//fmt.Println(age) returns no variable found
}
