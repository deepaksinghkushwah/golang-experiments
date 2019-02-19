package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	checkErr(err)
	win.SetTitle("Hello World")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l, err := gtk.LabelNew("Simple Label")
	checkErr(err)

	win.Add(l)

	win.SetDefaultSize(800, 600)

	win.ShowAll()
	gtk.Main()
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
