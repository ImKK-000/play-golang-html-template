package generate_test

import "testing"
import . "play-golang"

func Test_NewHTML_Should_Be_HTML(t *testing.T) {
	expectedHTML := `<h1>hello</h1>`
	htmlTemplate := `<h1>{{.Text}}</h1>`
	text := "hello"

	actualHTML := NewHTML(htmlTemplate, text)

	if expectedHTML != actualHTML {
		t.Errorf("expect %s but it got %s", expectedHTML, actualHTML)
	}
}
