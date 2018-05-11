package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
	fmt.Println("HelloWorld")
}
