package main

import "fmt"

func main() {
	//index always starts from 0
	var firstNames [3]string
	firstNames[0] = "naruto"
	firstNames[1] = "sasuke"
	firstNames[2] = "sakura"
	//names 3 could not be inserted because the limit is 3.

	//it is also possible to declare it directly like this.
	var lastNames = [3]string{
		"uzumaki", "uchiha", "haruno",
	}

	///printing values
	fmt.Println(len(firstNames)) //array length
	fmt.Println(lastNames[0])
	//changing the name it is possbile
	firstNames[0] = "boruto"
	fmt.Println(firstNames[0])
}
