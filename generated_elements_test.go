package htmlwriter

import (
	"testing"
)

// Sample tests for checking that elements are generated correctly

// Non-void element
func TestBody(t *testing.T) {
	writer.Body(writer.Attrs("class=foo"), func() {
		writer.Text("Foobar")
	})

	assertBufferMatches(t, "<body class=\"foo\">Foobar</body>", "body element was invalid.")
}

func TestBody_(t *testing.T) {
	writer.Body_(func() {
		writer.Text("hello world")
	})

	assertBufferMatches(t, "<body>hello world</body>", "body element without attributes was invalid.")
}

// Void element
func TestInput(t *testing.T) {
	writer.Input(writer.Attrs("type=text"))
	assertBufferMatches(t, "<input type=\"text\">", "input element was invalid.")
}

func TestInput_(t *testing.T) {
	writer.Input_()
	assertBufferMatches(t, "<input>", "input element without attributes was invalid.")
}

// Text only element
func TestTitle(t *testing.T) {
	writer.Title(writer.Attrs("lang=en"), "lorem ipsum")
	assertBufferMatches(t, "<title lang=\"en\">lorem ipsum</title>", "title element was invalid.")
}

func TestTitle_(t *testing.T) {
	writer.Title_("lorem ipsum dolor si amit")
	assertBufferMatches(t, "<title>lorem ipsum dolor si amit</title>", "title element was invalid.")
}

// Input element text
func TestSubmitInput(t *testing.T) {
	writer.SubmitInput(writer.Attrs("value=click"))
	assertBufferMatches(t, "<input value=\"click\" type=\"submit\">", "input element was invalid.")
}

func TestSubmitInput_(t *testing.T) {
	writer.SubmitInput_()
	assertBufferMatches(t, "<input type=\"submit\">", "input element without attributes was invalid.")
}
