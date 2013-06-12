package htmlbuffer

import (
	"bytes"
	"html"
	"io"
	"net/http"
)

type Attrs map[string]string

func NewHtmlBuffer() *HtmlBuffer {
	return &HtmlBuffer{}
}

type HtmlBuffer struct {
	bytes.Buffer
}

func WithHtmlBuffer(f func(b *HtmlBuffer)) string {
	html := NewHtmlBuffer()
	f(html)
	return html.String()
}

func (b *HtmlBuffer) Html5(attrs Attrs, innerHtml func()) {
	b.WriteString("<!DOCTYPE html>\n")
	b.WriteElement("html", attrs, innerHtml)
}

func (b *HtmlBuffer) Text(text string) {
	text = html.EscapeString(text)
	b.RawText(text)
}

func (b *HtmlBuffer) TextF(text string) func() {
	return func() {
		b.Text(text)
	}
}

func (b *HtmlBuffer) RawText(text string) {
	b.WriteString(text)
}

func (b *HtmlBuffer) RawTextF(text string) func() {
	return func() {
		b.RawText(text)
	}
}

func (b *HtmlBuffer) WriteToResponse(res http.ResponseWriter) {
	io.WriteString(res, b.String())
}

func (b *HtmlBuffer) WriteElement(tagName string, attrs Attrs, innerHtml func()) {
	b.WriteString("<" + tagName)
	b.writeAttributes(attrs)
	b.WriteString(">")

	if innerHtml != nil {
		innerHtml()
		b.WriteString("</" + tagName + ">")
	}
}

func (b *HtmlBuffer) WriteVoidElement(tagName string, attrs Attrs) {
	b.WriteElement(tagName, attrs, nil)
}

func (b *HtmlBuffer) writeAttributes(attrs Attrs) {
	if attrs != nil {
		for key, value := range attrs {
			b.WriteString(" " + key)

			if value != "" {
				b.WriteString("=\"" + value + "\"")
			}
		}
	}
}
