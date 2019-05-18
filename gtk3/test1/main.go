package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {

	gtk.Init(nil)
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	checkErr(err)
	win.SetTitle("My First Window")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//l, err := gtk.LabelNew("Hello World")
	//checkErr(err)

	//win.Add(l)

	b, err := gtk.ButtonNew()
	checkErr(err)
	b.SetLabel("Click Me")
	b.Connect("clicked", func() {
		fmt.Println("You clicked on button")
	})
	win.Add(b)

	win.SetDefaultSize(800, 600)
	win.ShowAll()
	gtk.Main()
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln("Error: ", e)
	}
}
