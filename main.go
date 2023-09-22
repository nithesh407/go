package main

import (
	"fmt"
	"go/util"
	// "net/http"
)
func main()  {
	greeting := "hello prasath"
	fmt.Println(greeting)
	fmt.Println(util.StringLength(greeting))
	// r := mux.NewRouter()
	// http.ListenAndServe(":9000",r)
}