package main

import "fmt"

func main()  {
	fmt.Printf("hello %T %v","nithesh","nithesh")
	fmt.Printf("\nhello %t",true)
	s := fmt.Sprintf("\nhello %s","palani") //stores the string with the formatter value to a new variable
	fmt.Println(s)
	fmt.Printf("\nhello %q %v","nithesh","nithesh") //%q for the quotation
	fmt.Printf("\nhello %9q %v","nithesh","nithesh") //%9q for the quotation and with nine wide
} 