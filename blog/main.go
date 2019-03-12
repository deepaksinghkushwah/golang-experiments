package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"

	"github.com/gorilla/mux"
)

var t *template.Template
var store *sessions.CookieStore

func init() {
	t = template.Must(template.ParseGlob("templates/*.html"))
	store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
	store.Options.HttpOnly = true
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/forbidden", forbiddenHandler).Methods("GET")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/test", testHandler).Methods("GET")
	r.HandleFunc("/logout", logoutHandler)

	r.HandleFunc("/", handler).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle string
	}{
		PageTitle: "Login Here",
	}
	t.ExecuteTemplate(w, "login.html", page)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/test", 301)
}
func handler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle string
		Name      string
		Age       int
	}{
		PageTitle: "Welcome to new world of golang",
		Name:      "Deepak Singh Kushwah",
		Age:       36,
	}
	t.ExecuteTemplate(w, "index.html", page)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle string
		Name      string
		Age       int
	}{
		PageTitle: "Hello World Handler",
		Name:      "Hello World",
		Age:       40,
	}
	t.ExecuteTemplate(w, "index.html", page)
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle string
		Name      string
	}{
		PageTitle: "403",
		Name:      "403",
	}
	t.ExecuteTemplate(w, "forbidden.html", page)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		http.Redirect(w, r, "/forbidden", 301)
		return
	}
	username, ok := untyped.(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/forbidden", 301)
		return
	}
	page := struct {
		PageTitle string
		Name      string
	}{
		PageTitle: "TEST Handler",
		Name:      username,
	}
	t.ExecuteTemplate(w, "test.html", page)
}

func loginCheck(r *http.Request) bool {
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		return false
	}
	_, ok = untyped.(string)
	if !ok {
		return false
	}
	return true
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	checkErr(err)
	session.Values["username"] = ""
	err = session.Save(r, w)
	checkErr(err)
	http.Redirect(w, r, "/", 301)
}

func checkErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
