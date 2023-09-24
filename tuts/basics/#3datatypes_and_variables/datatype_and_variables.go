package main

import "fmt"

func main()  {
	var num uint8 = 255
	var num2 uint16 = 260
	var number int16 = 440
	var name string = "hello nithesh"
	var implicit = 260 //implicitly guess the datatype
	n := 20 //expression assignment operator
	number += 5
	var x int //default value will be 0
	var bl bool //default value will be false
	bl = true //dafault value will be change to true
	fmt.Println(name)
	fmt.Println(number)
	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(implicit)
	fmt.Printf("%T",implicit)
	println()
	fmt.Printf("%T",n)
	println()
	fmt.Println(x,bl)

}