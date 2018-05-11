package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/leekchan/accounting"
)

func customFormat(t string) string {
	newT, _ := time.Parse("01-02-2011", t)
	return newT.Format("02-01-2006")
}

func currency(s float64) string {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac.FormatMoney(s)
}

var fm = template.FuncMap{
	"customFormat": customFormat,
	"curr":         currency,
}

type myData struct {
	Date   string
	Amount float64
}

func main() {
	data := myData{Date: "08-05-2018", Amount: 5000.55}

	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("time.html"))
	err := tpl.ExecuteTemplate(os.Stdout, "time.html", data)
	if err != nil {
		log.Fatalln(err)
	}
}
