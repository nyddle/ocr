package main

import (
	"fmt"
	"github.com/otiai10/gosseract"
)
//extract_text(filename, language="rus")

func main() {
	client := gosseract.NewClient()
	defer client.Close()

        client.Languages = []string{"rus"}

	client.SetImage("/Users/nyddle/Documents/testocr.jpg")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
