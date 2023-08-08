package maliciouspdfdetector

import (
	"os"
)

type PDFFile struct {
	Uri           string
	bytes         []byte
	keywordsCount map[string]int
}

// Initiate a new malicious detector from a PDF file URL
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
		return err
	}

	return nil
}

func (f *PDFFile) ParsePdfFile() {

	keywordsMap := make(map[string]keywordData)

	for _, b := range keywords {
		keywordsMap[string(b)] = keywordData{
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

type keywordData struct {
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
