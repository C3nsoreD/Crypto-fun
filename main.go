package main

import (
	"fmt"
	_ "math"
	"strings"
	"testing"
)

func main() {
	// y^2 = x^3 + ax + b
	// y^2 = x^3 + 7 --> used by bitcoin cryptography for private key

	// fName1 := "some.z"
	// fName2 := "some.zpl.exe"
	// f, err := changeLabelFileFormatToPDF(fName1)
	// f2, err := changeLabelFileFormatToPDF(fName2)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// // fmt.Print(f)
	// fmt.Print(f2)
}

func BenchmarkChangeLabelFileFormatToPDF(b *testing.B) {
	files := []string{"some.zpl.exe", "some.zpl"}

	for i := 0; i < b.N; i = i % len(files) {
		changeLabelFileFormatToPDF(files[i])
	}
}

func changeLabelFileFormatToPDF(fileName string) (string, error) {
	fileParts := strings.SplitAfter(fileName, ".")
	fileFormat := fileParts[len(fileParts)-1]
	switch fileFormat {
	case "zpl", "png", "svg":
		return strings.Replace(fileName, fileFormat, "pdf", 1), nil
	case "pdf":
		// For now it's only the case for .PDF format, and we do not need to replace it.
		return fileName, nil
	default:
		return "", fmt.Errorf("unsuporeted label format %q", fileFormat)

	}
}
