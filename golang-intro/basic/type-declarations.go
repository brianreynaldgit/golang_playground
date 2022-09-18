package main

import "fmt"

func main() {
	//to make it easier, we could create a new type and assign it as a variable.
	//for instance, say we want to create types to store personal informations.
	type firstName string
	type lastName string
	type age int

	var firstName1 firstName = "ash"
	var lastName1 lastName = "ketchum"
	var age1 age = 10
	var firstName2 firstName = "naruto"
	var lastName2 lastName = "uzumaki"
	var age2 age = 10

	fmt.Println(firstName1, lastName1, age1)
	fmt.Println(firstName2, lastName2, age2)

}
