package generate

import (
	"html/template"
)

type HTMLValiables struct {
	Text string
}

type writer struct {
	Data []byte
}

func (writer *writer) Write(bytes []byte) (int, error) {
	writer.Data = append(writer.Data, bytes...)
	return 0, nil
}

// HTML - pass html template and text string for generate new html template
func NewHTML(htmlTemplate, text string) string {
	htmlVariables := HTMLValiables{
		Text: text,
	}
	readTemplate, _ := template.New("newtemplate").Parse(htmlTemplate)
	writer := writer{}
	readTemplate.Execute(&writer, htmlVariables)
	return string(writer.Data)
}
