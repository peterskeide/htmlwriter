package htmlwriter

import (
	"fmt"
	"html"
	"io"
	"strings"
)

type Attrs map[string]string

func NewHtmlWriter(writer io.Writer) *HtmlWriter {
	return &HtmlWriter{writer}
}

type HtmlWriter struct {
	io.Writer
}

func (b *HtmlWriter) WriteString(str string) {
	b.Write([]byte(str))
}

func (b *HtmlWriter) Text(formatStr string, a ...interface{}) {
	text := html.EscapeString(fmt.Sprintf(formatStr, a...))
	b.WriteString(text)
}

func (b *HtmlWriter) TextF(formatStr string, a ...interface{}) func() {
	return func() {
		b.Text(formatStr, a...)
	}
}

func (b *HtmlWriter) RawText(formatStr string, a ...interface{}) {
	b.WriteString(fmt.Sprintf(formatStr, a...))
}

func (b *HtmlWriter) RawTextF(formatStr string, a ...interface{}) func() {
	return func() {
		b.RawText(formatStr, a...)
	}
}

func (b *HtmlWriter) WriteNormalElement(tagName string, attrs Attrs, innerHtml func()) {
	b.writeElement(tagName, attrs, innerHtml, true)
}

func (b *HtmlWriter) WriteVoidElement(tagName string, attrs Attrs) {
	b.writeElement(tagName, attrs, nil, false)
}

func (b *HtmlWriter) writeElement(tagName string, attrs Attrs, innerHtml func(), close bool) {
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

func (b *HtmlWriter) writeAttributes(attrs Attrs) {
	if attrs != nil {
		for key, value := range attrs {
			b.WriteString(" " + key)

			if value != "" {
				b.WriteString("=\"" + value + "\"")
			}
		}
	}
}

func (b *HtmlWriter) Attrs(formatStringAttrs string, a ...interface{}) Attrs {
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
