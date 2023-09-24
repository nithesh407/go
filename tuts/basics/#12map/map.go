package main

import "fmt"

func main() {
	mp := map[string]int{
		"apple":1,
		"orange":2,
		"pineapple":3,
	}
	fmt.Println(mp)

	fmt.Println(mp["apple"])

	mp["apple"]=230 //changing a value

	mp["banana"]=45 //adding new value
	
	fmt.Println(mp)

	delete(mp,"apple")

	fmt.Println(mp)
}