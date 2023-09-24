package main

import "fmt"

func main() {
	//slices
	//slices and arrays are different they are not same

	x := [5]int{1,2,3,4,5} //array declaration
	s := x[:] //slice declaration
	s = x[1:cap(x)]
	a := make([]int,5)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Printf("%T",a)
}