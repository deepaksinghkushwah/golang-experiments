package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func handler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	checkError(err)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
