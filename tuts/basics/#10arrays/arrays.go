package main

import "fmt"

func main() {
	// 1D array

	arr := [3]int{1,2,3}
	fmt.Println(arr)

	var sum int 
	for i := 0; i < len(arr); i++ {
		sum+=arr[i]
	}
	fmt.Println(sum)

	//2D array

	arr2 :=[2][3]int{{1,2,3},{3,4,5}}
	fmt.Println(arr2)
}