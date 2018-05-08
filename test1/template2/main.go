package main

import (
	"html/template"
	"os"
	"strings"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func main() {
	tpl := template.Must(template.New("Dee").Funcs(fm).Parse("uc Hello World"))
	tpl.Execute(os.Stdout, nil)
}
