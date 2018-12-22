package generate

import (
	"html/template"
	fakeWriter "play-golang/writer"
)

type HTMLValiables struct {
	Text string
}

// NewHTML - pass html template and html valiable struct for generate new html template
func NewHTML(htmlTemplate string, htmlValiables HTMLValiables) string {
	readTemplate, _ := template.New("newtemplate").Parse(htmlTemplate)
	writer := fakeWriter.Writer{}
	readTemplate.Execute(&writer, htmlValiables)
	return string(writer.Data)
}
