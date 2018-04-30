package app

import (
	"html/template"
	"log"
	"net/http"

	"github.com/deepaksinghkushwah/site1/models"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// Home page function for site
func Home(w http.ResponseWriter, r *http.Request) {
	page := models.Page{Title: "Home Page", Keywords: "Home Page", Description: "Home description"}
	err := tpl.ExecuteTemplate(w, "index.html", page)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// About function to show about us page
func About(w http.ResponseWriter, r *http.Request) {
	page := models.Page{Title: "About Us", Keywords: "About Page Keywords", Description: "About description"}
	err := tpl.ExecuteTemplate(w, "about.html", page)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// Contact us function
func Contact(w http.ResponseWriter, r *http.Request) {
	page := models.Page{Title: "Contact Us", Keywords: "Contact Page Keywords", Description: "Contact description"}
	err := tpl.ExecuteTemplate(w, "contact.html", page)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
