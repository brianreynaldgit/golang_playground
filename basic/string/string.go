package main

import "fmt"

func main() {
	/// String is basically
	var firstName string = "buzz fizz"
	fmt.Println(firstName)
	fmt.Println(len(firstName))

	//backticks is used for raw string. basically prints whatever you write there
	fmt.Println(`i do not know how to spell "pokemon" \n`)

	//it is possible to fetch a single character since character from the string.
	fmt.Printf("%v\n", firstName[0])
	//notice how it is printed as a number. because it is a number representation.
	fmt.Printf("%c\n", firstName[0])
}
