package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"play-golang/generate"
)

type HTMLVariables struct {
	Title string
	Names []string
}

func main() {
	readHTML, _ := ioutil.ReadFile("index.html")
	htmlVariables := HTMLVariables{
		Title: "Name of Persons",
		Names: []string{
			"Nattakit Boonyang",
			"Imkk 000",
		},
	}
	outputHTML := generate.NewHTML(string(readHTML), htmlVariables)
	fmt.Println(outputHTML)
	ioutil.WriteFile("output.html", []byte(outputHTML), os.ModePerm)
}
