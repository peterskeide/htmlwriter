package htmlwriter

func (b *HtmlWriter) Html5(attrs Attrs, innerHtml func()) {
	b.WriteString("<!DOCTYPE html>\n")
	b.Html(attrs, innerHtml)
}

func (b *HtmlWriter) Html5_(innerHtml func()) {
	b.Html5(nil, innerHtml)
}

func (b *HtmlWriter) StylesheetLink(href string) {
	b.Link(Attrs{"href": href, "type": "text/css", "rel": "stylesheet"})
}

func (b *HtmlWriter) ExternalScript(src string) {
	b.Script(Attrs{"src": src, "type": "text/javascript"}, nil)
}
