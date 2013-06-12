package htmlbuffer

func (b *HtmlBuffer) Base(attrs Attrs) {
	b.WriteVoidElement("base", attrs)
}

func (b *HtmlBuffer) Link(attrs Attrs) {
	b.WriteVoidElement("link", attrs)
}

func (b *HtmlBuffer) Meta(attrs Attrs) {
	b.WriteVoidElement("meta", attrs)
}

func (b *HtmlBuffer) Hr(attrs Attrs) {
	b.WriteVoidElement("hr", attrs)
}

func (b *HtmlBuffer) Br(attrs Attrs) {
	b.WriteVoidElement("br", attrs)
}

func (b *HtmlBuffer) Wbr(attrs Attrs) {
	b.WriteVoidElement("wbr", attrs)
}

func (b *HtmlBuffer) Input(attrs Attrs) {
	b.WriteVoidElement("input", attrs)
}

func (b *HtmlBuffer) Area(attrs Attrs) {
	b.WriteVoidElement("area", attrs)
}

func (b *HtmlBuffer) Command(attrs Attrs) {
	b.WriteVoidElement("command", attrs)
}

func (b *HtmlBuffer) Col(attrs Attrs) {
	b.WriteVoidElement("col", attrs)
}

func (b *HtmlBuffer) Img(attrs Attrs) {
	b.WriteVoidElement("img", attrs)
}

func (b *HtmlBuffer) Keygen(attrs Attrs) {
	b.WriteVoidElement("keygen", attrs)
}

func (b *HtmlBuffer) Param(attrs Attrs) {
	b.WriteVoidElement("param", attrs)
}

func (b *HtmlBuffer) Source(attrs Attrs) {
	b.WriteVoidElement("source", attrs)
}

func (b *HtmlBuffer) Track(attrs Attrs) {
	b.WriteVoidElement("track", attrs)
}

func (b *HtmlBuffer) Embed(attrs Attrs) {
	b.WriteVoidElement("embed", attrs)
}

