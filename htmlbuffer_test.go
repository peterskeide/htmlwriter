package htmlbuffer

import (
	"testing"
)

var buffer = NewHtmlBuffer()

func assertBufferMatches(t *testing.T, expected string, errorMessage string) {
	defer buffer.Reset()

	if res := buffer.String(); res != expected {
		t.Error(errorMessage+" Got:", res)
	}
}

// --- Tests for the HtmlBuffer type ---

// HtmlBuffer#WriteElement
func TestWriteElementWritesStartTag(t *testing.T) {
	buffer.WriteElement("input", nil, nil, false)
	assertBufferMatches(t, "<input>", "Failed to write valid start tag")
}

// HtmlBuffer#WriteElement
func TestWriteElementWritesStartTagWithAttributes(t *testing.T) {
	buffer.WriteElement("input", Attrs{"id": "test", "class": "main"}, nil, false)
	assertBufferMatches(t, "<input id=\"test\" class=\"main\">", "Failed to write valid start tag with attributes.")
}

// HtmlBuffer#WriteElement
func TestWriteElementShouldWriteElementWithInnerHtml(t *testing.T) {
	buffer.WriteElement("div", nil, func() {
		buffer.WriteElement("span", Attrs{"class": "highlight"}, buffer.TextF("lorem ipsum"), true)
	}, true)

	assertBufferMatches(t, "<div><span class=\"highlight\">lorem ipsum</span></div>", "Failed to write element with inner html.")
}

// HtmlBuffer#RawText
func TestRawTextWritesStringToBuffer(t *testing.T) {
	buffer.RawText("lorem ipsum")
	assertBufferMatches(t, "lorem ipsum", "Buffer did not contain expected raw text.")
}

// HtmlBuffer#RawTextF
func TestRawTextFReturnsFuncThatWritesStringToBuffer(t *testing.T) {
	innerHtml := buffer.RawTextF("lorem ipsum")
	innerHtml()
	assertBufferMatches(t, "lorem ipsum", "Buffer did not contain expected raw text.")
}

// HtmlBuffer#Text
func TestTextWritesStringToBuffer(t *testing.T) {
	buffer.Text("bilbo baggins")
	assertBufferMatches(t, "bilbo baggins", "Buffer did not contain expected text.")
}

// HtmlBuffer#Text
func TestTextWritesEscapedHtmlToBuffer(t *testing.T) {
	buffer.Text("<span>yo!</span>")
	assertBufferMatches(t, "&lt;span&gt;yo!&lt;/span&gt;", "Buffer did not contain escaped html.")
}

// HtmlBuffer#TextF
func TestTextFReturnsFuncThatWritesStringToBuffer(t *testing.T) {
	innerHtml := buffer.TextF("foobar")
	innerHtml()
	assertBufferMatches(t, "foobar", "Buffer did not contain expected text.")
}

// HtmlBuffer#TextF
func TestTextFReturnsFuncThatWritesEscapedHtmlToBuffer(t *testing.T) {
	innerHtml := buffer.TextF("<div>hello world</div>")
	innerHtml()
	assertBufferMatches(t, "&lt;div&gt;hello world&lt;/div&gt;", "Buffer did not contain escaped html.")
}

// HtmlBuffer#WriteNormalElement
func TestWriteNormalElementWritesElementWithStartAndEndTags(t *testing.T) {
	buffer.WriteNormalElement("a", nil, nil)
	assertBufferMatches(t, "<a></a>", "Failed to write element with start and end tags")
}

// HtmlBuffer#WriteNormalElement
func TestWriteNormalElementWritesElementWithAttributes(t *testing.T) {
	buffer.WriteNormalElement("a", Attrs{"href": "http://www.example.com"}, nil)
	assertBufferMatches(t, "<a href=\"http://www.example.com\"></a>", "Failed to write normal element with attributes")
}

// HtmlBuffer#WriteVoidElement
func TestWriteVoidElementWritestVoidElement(t *testing.T) {
	buffer.WriteVoidElement("link", nil)
	assertBufferMatches(t, "<link>", "Failed to write valid void element")
}

// HtmlBuffer#WriteVoidElement
func TestWriteVoidElementWritestVoidElementWithAttributes(t *testing.T) {
	buffer.WriteVoidElement("link", Attrs{"href": "stylesheets/main.css", "type": "text/css"})
	assertBufferMatches(t, "<link href=\"stylesheets/main.css\" type=\"text/css\">", "Failed to write valid void element with attributes")
}

// HtmlBuffer#Html5
func TestHtml5WritesDoctypeAndHtmlElement(t *testing.T) {
	buffer.Html5(nil, buffer.TextF(""))
	assertBufferMatches(t, "<!DOCTYPE html>\n<html></html>", "Failed to write valid html element with doctype")
}

// HtmlBuffer#Html5
func TestHtml5WritesDoctypeAndHtmlElementWithAttributes(t *testing.T) {
	buffer.Html5(Attrs{"lang": "en"}, buffer.TextF(""))
	assertBufferMatches(t, "<!DOCTYPE html>\n<html lang=\"en\"></html>", "Failed to write valid html element with attributes and doctype")
}

// HtmlBuffer#Html5_
func TestHtml5_WritesDoctypeAndHtmlElement(t *testing.T) {
	buffer.Html5_(buffer.TextF(""))
	assertBufferMatches(t, "<!DOCTYPE html>\n<html></html>", "Failed to write valid html element with doctype")
}
