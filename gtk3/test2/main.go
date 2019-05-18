package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)
	b, err := gtk.BuilderNew()
	checkErr(err)

	err = b.AddFromFile("../glade/myui.glade")
	checkErr(err)

	obj, err := b.GetObject("winMain")
	checkErr(err)

	win := obj.(*gtk.Window)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	login, err := b.GetObject("btnLogin")
	checkErr(err)

	btnLogin := login.(*gtk.Button)
	btnLogin.Connect("clicked", func() {
		username := getUsername(b)
		fmt.Println(username)

		password := getPassword(b)
		fmt.Println(password)
	})

	win.ShowAll()
	gtk.Main()
}

func getUsername(b *gtk.Builder) string {
	obj, err := b.GetObject("txtUsername")
	checkErr(err)
	s, _ := obj.(*gtk.Entry).GetText()
	return s
}

func getPassword(b *gtk.Builder) string {
	obj, err := b.GetObject("txtPassword")
	checkErr(err)
	s, _ := obj.(*gtk.Entry).GetText()
	return s
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln("Error: ", e)
	}
}
