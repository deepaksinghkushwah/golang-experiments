package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

type page struct {
	Name      string
	Gender    string
	Subscribe string
}

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("name")
		gender := r.FormValue("gender")
		subscribe := r.FormValue("subscribe")
		p := page{Name: name, Gender: gender}
		if subscribe == "on" {
			p.Subscribe = "Yes"
		} else {
			p.Subscribe = "No"
		}
		err := tpl.ExecuteTemplate(w, "index.html", p)
		checkError(err)
	} else {
		err := tpl.ExecuteTemplate(w, "index.html", nil)
		checkError(err)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}
