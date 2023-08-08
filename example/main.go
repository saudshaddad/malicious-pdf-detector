package main

import (
	"fmt"

	maliciouspdfdetector "github.com/saudshaddad/malicious-pdf-detector"
)

func main() {
	// initiate a new detector from the file URL
	PDFFile := maliciouspdfdetector.NewPDFFile("sample-mal.pdf")
	// PDFFile := maliciouspdfdetector.NewPDFFile("sample-clean.pdf")

	// read the file
	err := PDFFile.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	// call the parse method
	PDFFile.ParsePdfFile()

	// check the file if malicious or not
	if PDFFile.IsMalicious() {
		// the file is probably malicious, exit!
		fmt.Println("The file is probably malicious")
		return
	}

	fmt.Println("The file passed the test")
}
