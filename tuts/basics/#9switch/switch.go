package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("enter the input: ")
	scanner.Scan()
	input,_ := strconv.ParseInt(scanner.Text(),10,64)
	switch {
	case input==1:
		fmt.Println("one")
	case input==2:
		fmt.Println("two")
	case input==3:
		fmt.Println("three")
	default:
		fmt.Println("not a valid number")			
	}
}