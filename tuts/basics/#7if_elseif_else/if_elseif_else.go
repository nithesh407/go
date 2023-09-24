package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Tell me your Age: ")
	scanner.Scan()
	age,err := strconv.ParseInt(scanner.Text(),10,64)
	if err!= nil{
		fmt.Println(err)
	}
	if age >= 18{
		fmt.Println("you can ride alone")
	}else if age >= 14 {
		fmt.Println("you can ride with your parent")
	}else{
		fmt.Println("you cannot ride")
	}
}