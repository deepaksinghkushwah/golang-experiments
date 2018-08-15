package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, v)
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}
