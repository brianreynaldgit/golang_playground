package main

import "fmt"

func main() {
	var months = [...]string{
		"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec",
	}
	var slice1 = months[1:3]
	//how to read the slice is, the pointer starts from index 1 which is the 2nd position.
	//why the printed one ends on mar? it is because the length is 3-1 = 2.
	// with the length of 2, since index 1 got printed (feb), there is only a spot left. which is filled by mar.
	fmt.Println(slice1)

	///other func related to slice
	//find slice length
	fmt.Println(len(slice1))
	//find slice capacity
	//capacity works by substracting original array length with the index pointer. in this case it is 12 - 1
	fmt.Println(cap(slice1))
	//append data to the slice
	slice2 := append(slice1, "new month")
	fmt.Println(slice2)

	//be careful, when changing slice with appended slice, it has risk to change original slice.
	slice2[0] = "changed feb"
	fmt.Println(months)
	fmt.Println(slice1)
	fmt.Println(slice2)

	//you can make new slice as well
	newSlice := make([]string, 2, 5) //length 2, capacity 5
	newSlice[0] = "asd"
	newSlice[1] = "dsa"
	//newSlice[2] = "qwe" index out of range.
	fmt.Println(newSlice)

	//copy slice is possible
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
