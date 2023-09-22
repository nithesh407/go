package main

import (
	"fmt"
	"net/http"
)


func httpwriter(w http.ResponseWriter,r *http.Request)  {
	switch r.URL.Path{
	case "/":
		fmt.Fprintf(w,"hello from server")
	case "/dashboard":
		fmt.Fprint(w,"dashboard")
	case "/order":
		fmt.Fprint(w,"orders")
	default:
		fmt.Fprint(w,"big fat error")		
	}
	fmt.Printf("handling request with %s method",r.Method)
}
func htmltag(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","text/html")
	fmt.Fprint(w,"<h1>Hello</h1>")
}

func main(){
	http.HandleFunc("/",htmltag)
	http.ListenAndServe(":8080",nil)
}