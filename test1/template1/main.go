package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type person struct {
	Name   string
	Age    int32
	Gender string
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.html"))
}

func main() {
	//err := tpl.ExecuteTemplate(os.Stdout, "index2.html", "Deepak")
	//checkErr(err)

	//sage := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//sage := map[string]string{"Deepak": "Deepak Singh Kushwah", "Manish": "Manish Sharma", "Ashish": "Ashish Yadav"}
	//err := tpl.ExecuteTemplate(os.Stdout, "range.html", sage)
	//checkErr(err)

	p1 := person{Name: "Deepak", Age: 35, Gender: "Male"}
	p2 := person{Name: "Ashish", Age: 55, Gender: "Male"}
	p3 := person{Name: "Manish", Age: 45, Gender: "Male"}
	p4 := person{Name: "Ashok", Age: 34, Gender: "Male"}
	sage := []person{p1, p2, p3, p4}
	err := tpl.ExecuteTemplate(os.Stdout, "range.html", sage)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}

}
