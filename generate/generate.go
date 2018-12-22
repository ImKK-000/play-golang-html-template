package generate

import (
	"html/template"
	fakeWriter "play-golang/writer"
)

// HTMLValiables - pass variables to html template
type HTMLValiables struct {
	Text string
}

// NewHTML - pass html template and html valiable struct for generate new html template
func NewHTML(htmlTemplate string, htmlValiables interface{}) string {
	readTemplate, _ := template.New("htmltemplate").Parse(htmlTemplate)
	writer := fakeWriter.Writer{}
	readTemplate.Execute(&writer, htmlValiables)
	return string(writer.Data)
}
