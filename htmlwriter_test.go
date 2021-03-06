package htmlwriter

import (
	"bytes"
	"testing"
)

var buffer = new(bytes.Buffer)
var writer = NewHtmlWriter(buffer)

func assertBufferMatches(t *testing.T, expected string, errorMessage string) {
	defer buffer.Reset()

	if res := buffer.String(); res != expected {
		t.Error(errorMessage+" Got:", res)
	}
}

// --- Tests for the HtmlWriter type ---

// HtmlWriter#writeElement
func TestWriteElementWritesStartTag(t *testing.T) {
	writer.writeElement("input", nil, nil, false)
	assertBufferMatches(t, "<input>", "Failed to write valid start tag")
}

// HtmlWriter#WriteElement
func TestWriteElementWritesStartTagWithAttributes(t *testing.T) {
	writer.writeElement("input", Attrs{"id": "test", "class": "main"}, nil, false)
	assertBufferMatches(t, "<input id=\"test\" class=\"main\">", "Failed to write valid start tag with attributes.")
}

// HtmlWriter#WriteElement
func TestWriteElementShouldWriteElementWithInnerHtml(t *testing.T) {
	writer.writeElement("div", nil, func() {
		writer.writeElement("span", Attrs{"class": "highlight"}, writer.TextF("lorem ipsum"), true)
	}, true)

	assertBufferMatches(t, "<div><span class=\"highlight\">lorem ipsum</span></div>", "Failed to write element with inner html.")
}

// HtmlWriter#RawText
func TestRawTextWritesStringToBuffer(t *testing.T) {
	writer.RawText("<span>lorem ipsum</span>")
	assertBufferMatches(t, "<span>lorem ipsum</span>", "Buffer did not contain expected raw text.")
}

// HtmlWriter#Text
func TestRawTextSupportsFormatVerbs(t *testing.T) {
	writer.RawText("the %s is the %s", "bird", "word")
	assertBufferMatches(t, "the bird is the word", "Expected the bird to be the word.")
}

// HtmlWriter#RawTextF
func TestRawTextFReturnsFuncThatWritesStringToBuffer(t *testing.T) {
	innerHtml := writer.RawTextF("<div>lorem ipsum</div>")
	innerHtml()
	assertBufferMatches(t, "<div>lorem ipsum</div>", "Buffer did not contain expected raw text.")
}

// HtmlWriter#Text
func TestTextWritesStringToBuffer(t *testing.T) {
	writer.Text("bilbo baggins")
	assertBufferMatches(t, "bilbo baggins", "Buffer did not contain expected text.")
}

// HtmlWriter#Text
func TestTextWritesEscapedHtmlToBuffer(t *testing.T) {
	writer.Text("<span>yo!</span>")
	assertBufferMatches(t, "&lt;span&gt;yo!&lt;/span&gt;", "Buffer did not contain escaped html.")
}

// HtmlWriter#Text
func TestTextSupportsFormatVerbs(t *testing.T) {
	writer.Text("hi, my name is %s", "mud")
	assertBufferMatches(t, "hi, my name is mud", "Buffer did not contain expected string.")
}

// HtmlWriter#TextF
func TestTextFReturnsFuncThatWritesStringToBuffer(t *testing.T) {
	innerHtml := writer.TextF("foobar")
	innerHtml()
	assertBufferMatches(t, "foobar", "Buffer did not contain expected text.")
}

// HtmlWriter#TextF
func TestTextFReturnsFuncThatWritesEscapedHtmlToBuffer(t *testing.T) {
	innerHtml := writer.TextF("<div>hello world</div>")
	innerHtml()
	assertBufferMatches(t, "&lt;div&gt;hello world&lt;/div&gt;", "Buffer did not contain escaped html.")
}

// HtmlWriter#WriteNormalElement
func TestWriteNormalElementWritesElementWithStartAndEndTags(t *testing.T) {
	writer.WriteNormalElement("a", nil, nil)
	assertBufferMatches(t, "<a></a>", "Failed to write element with start and end tags")
}

// HtmlWriter#WriteNormalElement
func TestWriteNormalElementWritesElementWithAttributes(t *testing.T) {
	writer.WriteNormalElement("a", Attrs{"href": "http://www.example.com"}, nil)
	assertBufferMatches(t, "<a href=\"http://www.example.com\"></a>", "Failed to write normal element with attributes")
}

// HtmlWriter#WriteVoidElement
func TestWriteVoidElementWritestVoidElement(t *testing.T) {
	writer.WriteVoidElement("link", nil)
	assertBufferMatches(t, "<link>", "Failed to write valid void element")
}

// HtmlWriter#WriteVoidElement
func TestWriteVoidElementWritestVoidElementWithAttributes(t *testing.T) {
	writer.WriteVoidElement("link", Attrs{"href": "stylesheets/main.css", "type": "text/css"})
	assertBufferMatches(t, "<link href=\"stylesheets/main.css\" type=\"text/css\">", "Failed to write valid void element with attributes")
}

// HtmlWriter#Html5
func TestHtml5WritesDoctypeAndHtmlElement(t *testing.T) {
	writer.Html5(nil, writer.TextF(""))
	assertBufferMatches(t, "<!DOCTYPE html>\n<html></html>", "Failed to write valid html element with doctype")
}

// HtmlWriter#Html5
func TestHtml5WritesDoctypeAndHtmlElementWithAttributes(t *testing.T) {
	writer.Html5(Attrs{"lang": "en"}, writer.TextF(""))
	assertBufferMatches(t, "<!DOCTYPE html>\n<html lang=\"en\"></html>", "Failed to write valid html element with attributes and doctype")
}

// HtmlWriter#Html5_
func TestHtml5_WritesDoctypeAndHtmlElement(t *testing.T) {
	writer.Html5_(writer.TextF(""))
	assertBufferMatches(t, "<!DOCTYPE html>\n<html></html>", "Failed to write valid html element with doctype")
}

// HtmlWriter#Attrs
func TestAttrsReturnsAttrsStructBasedOnGivenStringAttributesAndValues(t *testing.T) {
	attrs := writer.Attrs("id=%s, class=main clearfix test, checked, data-test=%s", "foo", "test")
	assertAttributeValid(t, "id", "foo", attrs)
	assertAttributeValid(t, "class", "main clearfix test", attrs)
	assertAttributeValid(t, "checked", "", attrs)
	assertAttributeValid(t, "data-test", "test", attrs)
}

func assertAttributeValid(t *testing.T, attrName string, expectedValue string, attrs Attrs) {
	if attrs[attrName] != expectedValue {
		t.Error(attrName, "attribute missing or incorrect")
	}
}

// HtmlWriter#Attrs
func TestAttrsPanicsIfAttributesStringIsInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but everyone remained calm")
		}
	}()

	writer.Attrs("id=foo=bar")
}
