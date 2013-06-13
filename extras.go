package htmlbuffer

func (b *HtmlBuffer) Html5(attrs Attrs, innerHtml func()) {
	b.WriteString("<!DOCTYPE html>\n")
	b.Html(attrs, innerHtml)
}

func (b *HtmlBuffer) Html5_(innerHtml func()) {
	b.Html5(nil, innerHtml)
}

func (b *HtmlBuffer) StylesheetLink(href string) {
	b.Link(Attrs{"href": href, "type": "text/css", "rel": "stylesheet"})
}

func (b *HtmlBuffer) ExternalScript(src string) {
	b.Script(Attrs{"src": src, "type": "text/javascript"}, nil)
}
