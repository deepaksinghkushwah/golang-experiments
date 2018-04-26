package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

// Page struct for creatng structured page
type Page struct {
	Title       string
	Keywords    string
	Description string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "Home Page", Keywords: "Home Page", Description: "Home description"}
	err := tpl.ExecuteTemplate(w, "index.html", page)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "About Us", Keywords: "About Page Keywords", Description: "About description"}
	err := tpl.ExecuteTemplate(w, "about.html", page)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "Contact Us", Keywords: "Contact Page Keywords", Description: "Contact description"}
	err := tpl.ExecuteTemplate(w, "contact.html", page)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
