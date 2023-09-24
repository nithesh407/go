package main

import "fmt"

func main() {
	// comparision - < > <= >= == !=

	x := 5
	val := x < 5
	str1 := "Nithesh"
	str2 := "nithesh"
	isSame := str1 == str2

	fmt.Printf("%t %t\n",val,isSame)

	//logical operator &&-and ||-or !-not

	value := (true || false) && !false
	value2 := value && !!true // !!- double negative 
	fmt.Println(value,value2)

}