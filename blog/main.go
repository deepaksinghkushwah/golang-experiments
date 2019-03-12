package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/gorilla/mux"
)

var t *template.Template
var store *sessions.CookieStore

type User struct {
	Username      string
	Authenticated bool
}

func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)
	t = template.Must(template.ParseGlob("templates/*.html"))
	store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)
	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
	gob.Register(User{})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/test", testHandler).Methods("GET")
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/forbidden", forbiddenHandler)

	r.HandleFunc("/", handler).Methods("GET")

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle string
	}{
		PageTitle: "403 Forbidden",
	}
	t.ExecuteTemplate(w, "forbidden.html", page)
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

	user := &User{
		Username:      username,
		Authenticated: true,
	}

	session.Values["user"] = user

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

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	page := struct {
		PageTitle string
		Name      string
		Age       int
	}{
		PageTitle: "Good Bye Handler",
		Name:      "Good Bye",
		Age:       35,
	}
	t.ExecuteTemplate(w, "index.html", page)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	checkErr(err)

	user := getUser(session)
	fmt.Println(user)
	if auth := user.Authenticated; !auth {
		http.Redirect(w, r, "/forbidden", 301)
	}
	page := struct {
		PageTitle string
		Name      string
	}{
		PageTitle: "Good Bye Handler",
		Name:      "Loggedin",
	}
	t.ExecuteTemplate(w, "test.html", page)
}

func loginCheck(r *http.Request) bool {
	session, _ := store.Get(r, "session")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	} else {
		fmt.Println(auth, ok, session.Values["username"].(string))
		return true
	}
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	checkErr(err)
	session.Values["user"] = User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/test", 301)
}

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
