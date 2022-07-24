package main

import (
	"fmt"
	"fun/search"
	"strings"
)

func main() {
	v := []int{0, 2, 4, 6, 8, 12, 33}
	x := search.LinearSearch(12, v)
	fmt.Printf("Found at index %d value %d in array %v", x, 12, v)
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
