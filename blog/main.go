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

type loggedIn struct {
	isLoggedIn bool
}

func init() {
	t = template.Must(template.ParseGlob("templates/*.html"))
	store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
	store.Options.HttpOnly = true
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/forbidden", forbiddenHandler).Methods("GET")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/test", testHandler).Methods("GET")
	r.HandleFunc("/logout", logoutHandler).Methods("GET", "POST")

	r.HandleFunc("/", handler).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	l := loginCheck(r)
	page := struct {
		PageTitle string
		LoggedIn  bool
	}{
		PageTitle: "Login Here",
		LoggedIn:  l,
	}
	t.ExecuteTemplate(w, "login.html", page)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	session, _ := store.Get(r, "deepak-store")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/test", 301)
}
func handler(w http.ResponseWriter, r *http.Request) {
	l := loginCheck(r)
	page := struct {
		PageTitle string
		Name      string
		Age       int
		LoggedIn  bool
	}{
		PageTitle: "Welcome to new world of golang",
		Name:      "Deepak Singh Kushwah",
		Age:       36,
		LoggedIn:  l,
	}
	t.ExecuteTemplate(w, "index.html", page)
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	l := loginCheck(r)
	page := struct {
		PageTitle string
		Name      string
		LoggedIn  bool
	}{
		PageTitle: "403",
		Name:      "403",
		LoggedIn:  l,
	}
	t.ExecuteTemplate(w, "forbidden.html", page)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	l := loginCheck(r)
	session, _ := store.Get(r, "deepak-store")
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
		LoggedIn  bool
	}{
		PageTitle: "TEST Handler",
		Name:      username,
		LoggedIn:  l,
	}
	t.ExecuteTemplate(w, "test.html", page)
}

func loginCheck(r *http.Request) bool {
	session, _ := store.Get(r, "deepak-store")
	untyped, ok := session.Values["username"]
	if !ok {
		return false
	}
	username, ok := untyped.(string)
	if username == "" {
		return false
	}
	return true
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "deepak-store")
	checkErr(err)
	session.Values["deepak-store"] = ""
	session.Values["username"] = ""
	err = session.Save(r, w)
	checkErr(err)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
