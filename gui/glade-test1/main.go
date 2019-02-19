package main

import (
	"errors"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"

	"github.com/gotk3/gotk3/gtk"
)

const appID = "com.deepak.gui.glade-test"
var txtUsername 

func main() {
	application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	checkErr(err)

	application.Connect("startup", func() {
		log.Println("Application startup initiated")
	})

	application.Connect("activate", func() {
		log.Println("Application activated")

		builder, err := gtk.BuilderNewFromFile("gui.glade")
		checkErr(err)
		signals := map[string]interface{}{
			"on_winMain_destroy":  onMainWindowDestroy,
			"on_btnLogin_clicked": onBtnLoginClicked,
		}
		builder.ConnectSignals(signals)

		obj, err := builder.GetObject("winMain")
		checkErr(err)

		win, err := isWindow(obj)
		checkErr(err)

		win.Show()
		application.AddWindow(win)
	})

	application.Connect("shuddown", func() {
		log.Println("Application shutdown")
	})

	os.Exit(application.Run(os.Args))

}

func onMainWindowDestroy() {
	log.Println("Main window destroyed")
}

func onBtnLoginClicked() {
	log.Println("Login button clicked")
}

func isWindow(obj glib.IObject) (*gtk.Window, error) {
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
