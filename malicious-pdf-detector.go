package maliciouspdfdetector

// package main

import (
	"fmt"
	"os"
)

type PDFFile struct {
	Uri           string
	bytes         []byte
	keywordsCount map[string]int
}

func NewPDFFile(uri string) *PDFFile {
	return &PDFFile{
		Uri:           uri,
		keywordsCount: make(map[string]int),
	}
}

func (f *PDFFile) ReadFile() error {
	var err error
	f.bytes, err = os.ReadFile(f.Uri)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (f *PDFFile) ParsePdfFile() {

	keywordsMap := make(map[string]KeywordData)

	for _, b := range keywords {
		keywordsMap[string(b)] = KeywordData{
			count:           0,
			length:          len(b),
			currentProgress: 0,
			bytes:           b,
		}
	}

	for _, b := range f.bytes {
		for key, value := range keywordsMap {
			if b == value.bytes[value.currentProgress] {
				value.currentProgress++
			} else {
				value.currentProgress = 0
				keywordsMap[key] = value
				continue
			}

			if value.currentProgress == len(value.bytes) {
				value.count++
				value.currentProgress = 0
				keywordsMap[key] = value
			} else {
				keywordsMap[key] = value
			}
		}
	}

	for key, value := range keywordsMap {
		f.keywordsCount[key] = value.count
	}
}

func (f *PDFFile) IsMalicious() bool {
	for key, count := range f.keywordsCount {
		if count > 0 {
			switch key {
			case "JS":
				return true
			case "JavaScript":
				return true
			case "/AA":
				return true
			case "/OpenAction":
				return true
			}
		}
	}

	return false
}

type KeywordData struct {
	count           int
	length          int
	currentProgress int
	bytes           []byte
}

var keywords []([]byte) = []([]byte){
	[]byte("obj"),
	[]byte("stream"),
	[]byte("xref"),
	[]byte("trailer"),
	[]byte("startxref"),
	[]byte("/Page"),
	[]byte("/Encrypt"),
	[]byte("/ObjStm"),
	[]byte("/JS"),
	[]byte("/JavaScript"),
	[]byte("/AA"),
	[]byte("/OpenAction"),
	[]byte("/AcroForm"),
	[]byte("/JBIG2Decode"),
	[]byte("/RichMedia"),
	[]byte("/Launch"),
	[]byte("/EmbeddedFile"),
	[]byte("/XFA"),
	[]byte("/Colors > 2^24"),
}

// func (f *PDFFile) passThrough() {

// 	objBytesCount := 0
// 	// probablyInsideObject := false
// 	isInsideObject := false

// 	for pos, b := range f.bytes {

// 		if isInsideObject {
// 			fmt.Println(b, f.bytes[pos+1])
// 			isInsideObject = false
// 		}

// 		switch b {
// 		case 'o':
// 			if objBytesCount == 0 {
// 				objBytesCount++
// 			} else {
// 				objBytesCount = 0
// 			}
// 		case 'b':
// 			if objBytesCount == 1 {
// 				objBytesCount++
// 			} else {
// 				objBytesCount = 0
// 			}
// 		case 'j':
// 			if objBytesCount == 2 {
// 				objBytesCount = 0
// 				// probablyInsideObject = true
// 			}
// 		default:
// 			objBytesCount = 0
// 		}
// 	}
// }

// func (f *PDFFile) FindStreams() {
// 	var reg = regexp.MustCompile(`stream\n(.*\n)+?endstream`)
// 	// var reg = regexp.MustCompile(`stream(.*\n).+?(?=\nendstream)`)

// 	// (?<=stream\n)(.*\n).+?(?=\nendstream)

// 	found := reg.FindAll(f.bytes, -1)

// 	fmt.Println("found: ", len(found))
// 	for index, by := range found {
// 		// convert byte slice to io.Reader
// 		sanitizedBytes := by[7 : len(by)-11]

// 		reader := bytes.NewReader(sanitizedBytes)

// 		r, err := zlib.NewReader(reader)
// 		if err != nil {
// 			fmt.Println("---------------------------------- ", index+1)
// 			fmt.Println(err)
// 			continue
// 		}
// 		fmt.Println("---------------------------------- ", index+1)
// 		io.Copy(os.Stdout, r)
// 		r.Close()
// 	}
// }
