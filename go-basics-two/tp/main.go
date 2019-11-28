package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	Td := Todo{}
	dataset := []Todo{}

	Td = Todo{Title: "Task 1", Done: false}
	dataset = append(dataset, Td)
	Td = Todo{Title: "Task 2", Done: true}
	dataset = append(dataset, Td)

	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos:     dataset,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
