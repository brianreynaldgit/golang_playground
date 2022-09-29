package main

import "fmt"

const (
	i = 7
	j
	k
)

func main() {
	// var s1 []int
	// s2 := []int{1, 2, 3}
	// n1 := copy(s1, s2)
	// fmt.Printf("%d,%v,%v\n", n1, s1, s2)
	// fmt.Println(s1 == nil)
	// a := []string{"A", "B", "C", "D", "E"}
	// a = a[:0]
	// fmt.Println(a, len(a), cap(a))

	// m := make(map[int]*int)

	// for i := 0; i < 3; i++ {
	// 	m[i] = &i
	// }
	// for _, v := range m {
	// 	print(*v)
	// }

	// fmt.Println("A")
	// c := make(chan string)
	// c <- "john"
	// fmt.Println("B")

	// x := interface{}(&S{"a", "b", "c"})
	// y := interface{}(&S{"a", "b", "c"})
	// fmt.Println(x == y)
	// s := "123"
	// ps := &s
	// b := []byte(*ps)
	// pb := &b

	// s += "4"
	// *ps += "5"
	// b[1] = '0'
	// print(*ps)
	// print(string(*pb))
	// var s *int
	// fmt.Println(s)

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	makeSquares(a)
	fmt.Println(a)
}

func makeSquares(array [10]int) {
	for index, elem := range array {
		array[index] = elem * elem
	}
}

type S struct {
	name string
}
