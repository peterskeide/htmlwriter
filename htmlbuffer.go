package htmlbuffer

import (
	"bytes"
	"html"
	"io"
	"net/http"
	"strings"
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
	b.WriteNormalElement("html", attrs, innerHtml)
}

func (b *HtmlBuffer) Html5_(innerHtml func()) {
	b.Html5(nil, innerHtml)
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

func (b *HtmlBuffer) WriteElement(tagName string, attrs Attrs, innerHtml func(), close bool) {
	b.WriteString("<" + tagName)
	b.writeAttributes(attrs)
	b.WriteString(">")

	if innerHtml != nil {
		innerHtml()
	}

	if close {
		b.WriteString("</" + tagName + ">")
	}
}

func (b *HtmlBuffer) WriteNormalElement(tagName string, attrs Attrs, innerHtml func()) {
	b.WriteElement(tagName, attrs, innerHtml, true)
}

func (b *HtmlBuffer) WriteVoidElement(tagName string, attrs Attrs) {
	b.WriteElement(tagName, attrs, nil, false)
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

func (b *HtmlBuffer) Attrs(strAttrs ...string) Attrs {
	attrs := Attrs{}

	for _, attr := range strAttrs {
		parts := strings.Split(attr, "=")
		switch len(parts) {
		case 1:
			attrs[parts[0]] = ""
		case 2:
			attrs[parts[0]] = parts[1]
		default:
			panic("Invalid attribute string: " + attr)
		}
	}

	return attrs
}
