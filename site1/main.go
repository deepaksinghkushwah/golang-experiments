package main

import (
	"net/http"

	"github.com/deepaksinghkushwah/site1/app"
)

func main() {

	//fs := http.FileServer(http.Dir("static"))

	http.HandleFunc("/", app.Home)
	http.HandleFunc("/about", app.About)
	http.HandleFunc("/contact", app.Contact)
	http.Handle("favicon.ico", http.NotFoundHandler())

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
