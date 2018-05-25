package ocr

import (
	"github.com/otiai10/gosseract"
)

func ExtractText(filename string, languages []string) (string, error) {
    client := gosseract.NewClient()
    defer client.Close()

    client.Languages = languages

    client.SetImage(filename)
    return client.Text()
}
