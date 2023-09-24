package main

import "fmt"

func main() {
	x := 1

	for i := 0; i < x; i++ {
		fmt.Println(i)
	}
	println()
	// for x <= 10{
	// 	fmt.Println(x)
	// 	if x%2==0{  //continue
	// 		fmt.Printf("%v is even",x)
	// 		continue //infinite loop
	// 	}
	// 	x++
	// }
	for i := 0; i < 6; i++ {
		fmt.Println(i)
		if i==5 {
			break //break statement
		}
	}

	
}