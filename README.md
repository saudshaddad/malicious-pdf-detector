# malicious-pdf-detector

This package is for detecting malicious PDF files. This package is still under development but eventually it can will be used to detect malicious PDF files uploads in a web server and probably to sanitize the file before saving. The package tries to detect embedded Javascript in the PDF they may be harmful.

**Note: This package is still under development and cannot detect all cases of evasion techniques. Use at your discretion.**

### Installation

To install the package use the following command

```sh
$ go get github.com/saudshaddad/malicious-pdf-detector
```

### Features

**The package tries to detect the following PDF Objects and Dictionaries:**
- ``/JS`` & ``/JavaScript``: objects used to execute javascript in the PDF file
- ``/AA`` & ``/OpenAction``:  objects used to execute action upon opening a PDF file.

There Dictionaries is not necessary harmful but in most cases it cause harm and is used by hackers to embed code to attack a machine.

**Not Implemented Objects:** ``/stream``  streams may contain additional malicious code

### Getting Started

The detector can be used as mentioned in the following example. You can also refer to /example that contains two PDF files

```go
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
	  
	// call the parse method, here most of the work is done
	PDFFile.ParsePdfFile()
	
	// check the file if malicious or not
	if PDFFile.IsMalicious() {
		// the file is probably malicious, exit!
		fmt.Println("The file is probably malicious")
		return
	}
	
	fmt.Println("The file passed the test")
}
```

