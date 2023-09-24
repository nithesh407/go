package main

import "fmt"

func main() {
	//range in slices

	a := []int{1,2,3,4,5,6}

	for i,ele := range a{ 
		fmt.Printf("%d : %d\n",i,ele) // i represents the index and ele represents the value of the slice
	}

	for _,ele := range a{
		fmt.Printf("%d\n",ele)
	}
}	