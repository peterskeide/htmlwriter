package htmlbuffer

import (
	"bytes"
	"fmt"
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

func (b *HtmlBuffer) WriteNormalElement(tagName string, attrs Attrs, innerHtml func()) {
	b.writeElement(tagName, attrs, innerHtml, true)
}

func (b *HtmlBuffer) WriteVoidElement(tagName string, attrs Attrs) {
	b.writeElement(tagName, attrs, nil, false)
}

func (b *HtmlBuffer) writeElement(tagName string, attrs Attrs, innerHtml func(), close bool) {
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

func (b *HtmlBuffer) Attrs(formatStringAttrs string, a ...interface{}) Attrs {
	attrs := Attrs{}

	formattedAttrs := fmt.Sprintf(formatStringAttrs, a...)
	parts := strings.Split(formattedAttrs, ",")

	for _, attr := range parts {
		attrAndValue := strings.Split(attr, "=")
		switch len(attrAndValue) {
		case 1:
			attrs[strings.TrimSpace(attrAndValue[0])] = ""
		case 2:
			attrs[strings.TrimSpace(attrAndValue[0])] = strings.TrimSpace(attrAndValue[1])
		default:
			panic("Invalid attribute declaration: " + attr)
		}
	}

	return attrs
}
