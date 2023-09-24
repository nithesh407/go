package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //creates a scanner input
	fmt.Printf("Enter your Born year : ")
	scanner.Scan() //scans the scanner object
	input,_ := strconv.ParseInt(scanner.Text(),10,64) //converting the scanner string to parseInt with th base 10 and 64
	fmt.Printf("your age will be %d in the end of 2023\n",2023-input)
}