package main

import (
	"fmt"
	"ocr/ocr"
	"code.sajari.com/docconv"
	"log"
)


func extractText(filename string) (string, error) {
	return "", nil
}

func main() {

	/*
	Надо сделаю модулек, который из файлов типов: doc, docx, xls, xlsx, rtf, html, pdf (текстовых), pdf (картиночных), jpg, png, bmp, tiff - извлекает текстовые данные
	 */

	for _, filename := range([]string{"/Users/nyddle/Downloads/alexanderdavydovcv.docx"}) {
		languages := []string{"rus"}
		text, err := ocr.ExtractText(filename, languages)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
	}

	res, err := docconv.ConvertPath("/Users/nyddle/Downloads/alexanderdavydovcv.docx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
