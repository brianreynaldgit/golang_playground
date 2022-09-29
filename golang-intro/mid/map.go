package main

import "fmt"

func main() {
	//map works on key value
	person := map[string]string{
		"firstname": "james",
		"lastname":  "bond",
	}
	fmt.Println(person)

	//to access specific data
	fmt.Println(person["firstname"])

	//other operations
	fmt.Println(len(person))
	book := make(map[string]string)
	book["title"] = "good ridance"
	book["title"] = "replaced"
	fmt.Println(book)
}
