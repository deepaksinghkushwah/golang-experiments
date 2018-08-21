package main

import (
	"html/template"
	"net/http"

	app "github.com/deepaksinghkushwah/site2/app"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", app.Home)
	http.HandleFunc("/features", app.Features)
	http.HandleFunc("/pricing", app.Pricing)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", nil)
}
