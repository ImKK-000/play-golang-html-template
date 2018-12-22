package main

import (
	"fmt"
	"play-golang/generate"
)

type HTMLVariables struct {
	Title string
	Names []string
}

const htmlTemplate = `<!DOCTYPE html><html lang="en"><body><h1>{{.Title}}</h1><ol>{{range .Names}}<li>{{.}}</li>{{else}}{{end}}</ol></body></html>`

func main() {
	htmlVariables := HTMLVariables{
		Title: "Name of Persons",
		Names: []string{
			"Nattakit Boonyang",
			"Imkk 000",
		},
	}
	outputHTML := generate.NewHTML(htmlTemplate, htmlVariables)
	fmt.Println(outputHTML)
}
