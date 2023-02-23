package main

import (
	"fmt"

	maliciouspdfdetector "github.com/saudshaddad/malicious-pdf-detector"
)

func main() {
	PDFFile := maliciouspdfdetector.NewPDFFile("sample.pdf")

	err := PDFFile.ReadFile()
	if err != nil {
		fmt.Println(err)
	}

	PDFFile.ParsePdfFile()
	fmt.Println(PDFFile.IsMalicious())
}
