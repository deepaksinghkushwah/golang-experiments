package app

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle   string
		PageHeading string
	}{PageTitle: "Home", PageHeading: "Welcome To Home Of GO Web Application Development"}
	err := tpl.ExecuteTemplate(w, "home.html", page)
	checkError(err)
}

func Features(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle   string
		PageHeading string
	}{PageTitle: "Features", PageHeading: "Features"}
	err := tpl.ExecuteTemplate(w, "features.html", page)
	checkError(err)
}

func Pricing(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle   string
		PageHeading string
	}{PageTitle: "Pricing", PageHeading: "Pricing"}
	err := tpl.ExecuteTemplate(w, "pricing.html", page)
	checkError(err)
}
