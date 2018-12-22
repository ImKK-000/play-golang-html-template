package generate_test

import (
	. "play-golang/generate"
	"testing"
)

func Test_NewHTML_Should_Be_HTML(t *testing.T) {
	expectedHTML := `<h1>hello</h1>`
	htmlTemplate := `<h1>{{.Text}}</h1>`
	htmlValiables := HTMLValiables{
		Text: "hello",
	}

	actualHTML := NewHTML(htmlTemplate, htmlValiables)

	if expectedHTML != actualHTML {
		t.Errorf("expect %s but it got %s", expectedHTML, actualHTML)
	}
}
