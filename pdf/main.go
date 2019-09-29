package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {
	fmt.Println("Hello World")
	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello World")
	err := pdf.OutputFileAndClose("mypdf.pdf")
	if err != nil {
		log.Fatalln(err)
	}
}
