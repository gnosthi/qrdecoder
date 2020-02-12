package main

import (
	"bytes"
	"fmt"
	"github.com/liyue201/goqr"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
)

// File contains an instance of a file.
type File struct {
	Location string
	FileData []byte
	DecodedData []*goqr.QRData
}

// DecodeQR decodes a QRCode
func DecodeQR(fileName string) (f File, e error) {
	f.Location = fileName
	fileData, ioErr := ioutil.ReadFile(f.Location)
	if ioErr != nil {
		return f, ioErr
	}
	f.FileData = fileData

	imgData, _, imgErr := image.Decode(bytes.NewReader(f.FileData))
	if imgErr != nil {
		return f, imgErr
	}

	decodedData, decodeErr := goqr.Recognize(imgData)
	if decodeErr != nil {
		return f, decodeErr
	}

	f.DecodedData = decodedData

	return f, nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Please specify a file")
	}


	for i := 1; i < len(os.Args); i++ {
		fileName := os.Args[i]
		fInfo, fErr := os.Stat(fileName)
		if os.IsNotExist(fErr) {
			fmt.Printf("Could not open file %s, file does not exist\n", fileName)
			continue
		}
		if fInfo.IsDir() {
			fmt.Printf("%s is a directory, which are not currently supported\n", fileName)
			continue
		}
		file, fileErr := DecodeQR(fileName)
		if fileErr != nil {
			fmt.Printf("An error occurred. %e", fileErr)
			os.Exit(1)
		}

		for _, data := range file.DecodedData {
			fmt.Printf("File: %s\n %s\n", fileName, data.Payload)
		}
	}
}
