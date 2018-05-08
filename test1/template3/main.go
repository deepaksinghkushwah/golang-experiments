package main

import (
	"html/template"
	"os"
	"time"
)

func customFormat(t time.Time) string {
	return t.Format("02-01-2006")
}

func currency(s string) string {
	return "$" + s
}

var fm = template.FuncMap{
	"customFormat": customFormat,
	"curr":         currency,
}

type myData struct {
	date   time.Time
	amount float64
}

func main() {
	t := time.Now()
	data := myData{date: t, amount: 5000.00}

	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
	tpl.ExecuteTemplate(os.Stdout, "time.html", data)
}
