package htmlbuffer

func (b *HtmlBuffer) StylesheetLink(href string) {
	b.Link(Attrs{"href": href, "type": "text/css", "rel": "stylesheet"})
}

func (b *HtmlBuffer) ExternalScript(src string) {
	b.Script(Attrs{"src": src, "type": "text/javascript"}, nil)
}
