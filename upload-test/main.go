package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

// Page struct which describe page structure
type Page struct {
	Title       string
	Keywords    string
	Description string
	Message     string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "Home Page", Keywords: "This is home page keywords", Description: "This is description of page"}
	err := tpl.ExecuteTemplate(w, "index.html", page)
	checkWebError(w, err)
}

func uploadhandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handle, err := r.FormFile("image")
		checkError(err)

		// we will use this id for new unique name
		id, err := uuid.NewV4()

		// extension will come with . (dot) extension
		ext := strings.Index(handle.Filename, ".")

		newname := fmt.Sprintf("%s", id) + handle.Filename[ext:]

		defer file.Close()
		f, err := os.OpenFile("./images/"+newname, os.O_WRONLY|os.O_CREATE, 0777)
		checkError(err)

		defer f.Close()
		io.Copy(f, file)
	}
	page := Page{Title: "File Uploaded", Keywords: "This is home page keywords", Description: "This is description of page", Message: "File Uploaded"}
	err := tpl.ExecuteTemplate(w, "index.html", page)
	checkWebError(w, err)
}

func checkWebError(w http.ResponseWriter, error error) {
	if error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func checkError(error error) {
	if error != nil {
		log.Println(error)
		return
	}
}

func main() {
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/upload", uploadhandler)
	http.ListenAndServe(":8080", nil)
}
