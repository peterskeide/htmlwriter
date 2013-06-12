package htmlbuffer

func (b *HtmlBuffer) Head(attrs Attrs, innerHtml func()) {
	b.WriteElement("head", attrs, innerHtml)
}

func (b *HtmlBuffer) Title(attrs Attrs, innerHtml func()) {
	b.WriteElement("title", attrs, innerHtml)
}

func (b *HtmlBuffer) Style(attrs Attrs, innerHtml func()) {
	b.WriteElement("style", attrs, innerHtml)
}

func (b *HtmlBuffer) Script(attrs Attrs, innerHtml func()) {
	b.WriteElement("script", attrs, innerHtml)
}

func (b *HtmlBuffer) Noscript(attrs Attrs, innerHtml func()) {
	b.WriteElement("noscript", attrs, innerHtml)
}

func (b *HtmlBuffer) Body(attrs Attrs, innerHtml func()) {
	b.WriteElement("body", attrs, innerHtml)
}

func (b *HtmlBuffer) Section(attrs Attrs, innerHtml func()) {
	b.WriteElement("section", attrs, innerHtml)
}

func (b *HtmlBuffer) Nav(attrs Attrs, innerHtml func()) {
	b.WriteElement("nav", attrs, innerHtml)
}

func (b *HtmlBuffer) Article(attrs Attrs, innerHtml func()) {
	b.WriteElement("article", attrs, innerHtml)
}

func (b *HtmlBuffer) Aside(attrs Attrs, innerHtml func()) {
	b.WriteElement("aside", attrs, innerHtml)
}

func (b *HtmlBuffer) H1(attrs Attrs, innerHtml func()) {
	b.WriteElement("h1", attrs, innerHtml)
}

func (b *HtmlBuffer) H2(attrs Attrs, innerHtml func()) {
	b.WriteElement("h2", attrs, innerHtml)
}

func (b *HtmlBuffer) H3(attrs Attrs, innerHtml func()) {
	b.WriteElement("h3", attrs, innerHtml)
}

func (b *HtmlBuffer) H4(attrs Attrs, innerHtml func()) {
	b.WriteElement("h4", attrs, innerHtml)
}

func (b *HtmlBuffer) H5(attrs Attrs, innerHtml func()) {
	b.WriteElement("h5", attrs, innerHtml)
}

func (b *HtmlBuffer) H6(attrs Attrs, innerHtml func()) {
	b.WriteElement("h6", attrs, innerHtml)
}

func (b *HtmlBuffer) Header(attrs Attrs, innerHtml func()) {
	b.WriteElement("header", attrs, innerHtml)
}

func (b *HtmlBuffer) Footer(attrs Attrs, innerHtml func()) {
	b.WriteElement("footer", attrs, innerHtml)
}

func (b *HtmlBuffer) Address(attrs Attrs, innerHtml func()) {
	b.WriteElement("address", attrs, innerHtml)
}

func (b *HtmlBuffer) Main(attrs Attrs, innerHtml func()) {
	b.WriteElement("main", attrs, innerHtml)
}

func (b *HtmlBuffer) P(attrs Attrs, innerHtml func()) {
	b.WriteElement("p", attrs, innerHtml)
}

func (b *HtmlBuffer) Pre(attrs Attrs, innerHtml func()) {
	b.WriteElement("pre", attrs, innerHtml)
}

func (b *HtmlBuffer) Blockquote(attrs Attrs, innerHtml func()) {
	b.WriteElement("blockquote", attrs, innerHtml)
}

func (b *HtmlBuffer) Ol(attrs Attrs, innerHtml func()) {
	b.WriteElement("ol", attrs, innerHtml)
}

func (b *HtmlBuffer) Ul(attrs Attrs, innerHtml func()) {
	b.WriteElement("ul", attrs, innerHtml)
}

func (b *HtmlBuffer) Li(attrs Attrs, innerHtml func()) {
	b.WriteElement("li", attrs, innerHtml)
}

func (b *HtmlBuffer) Dl(attrs Attrs, innerHtml func()) {
	b.WriteElement("dl", attrs, innerHtml)
}

func (b *HtmlBuffer) Dt(attrs Attrs, innerHtml func()) {
	b.WriteElement("dt", attrs, innerHtml)
}

func (b *HtmlBuffer) Dd(attrs Attrs, innerHtml func()) {
	b.WriteElement("dd", attrs, innerHtml)
}

func (b *HtmlBuffer) Figure(attrs Attrs, innerHtml func()) {
	b.WriteElement("figure", attrs, innerHtml)
}

func (b *HtmlBuffer) Figcaption(attrs Attrs, innerHtml func()) {
	b.WriteElement("figcaption", attrs, innerHtml)
}

func (b *HtmlBuffer) Div(attrs Attrs, innerHtml func()) {
	b.WriteElement("div", attrs, innerHtml)
}

func (b *HtmlBuffer) A(attrs Attrs, innerHtml func()) {
	b.WriteElement("a", attrs, innerHtml)
}

func (b *HtmlBuffer) Em(attrs Attrs, innerHtml func()) {
	b.WriteElement("em", attrs, innerHtml)
}

func (b *HtmlBuffer) Strong(attrs Attrs, innerHtml func()) {
	b.WriteElement("strong", attrs, innerHtml)
}

func (b *HtmlBuffer) Small(attrs Attrs, innerHtml func()) {
	b.WriteElement("small", attrs, innerHtml)
}

func (b *HtmlBuffer) S(attrs Attrs, innerHtml func()) {
	b.WriteElement("s", attrs, innerHtml)
}

func (b *HtmlBuffer) Cite(attrs Attrs, innerHtml func()) {
	b.WriteElement("cite", attrs, innerHtml)
}

func (b *HtmlBuffer) Q(attrs Attrs, innerHtml func()) {
	b.WriteElement("q", attrs, innerHtml)
}

func (b *HtmlBuffer) Dfn(attrs Attrs, innerHtml func()) {
	b.WriteElement("dfn", attrs, innerHtml)
}

func (b *HtmlBuffer) Abbr(attrs Attrs, innerHtml func()) {
	b.WriteElement("abbr", attrs, innerHtml)
}

func (b *HtmlBuffer) Data(attrs Attrs, innerHtml func()) {
	b.WriteElement("data", attrs, innerHtml)
}

func (b *HtmlBuffer) Time(attrs Attrs, innerHtml func()) {
	b.WriteElement("time", attrs, innerHtml)
}

func (b *HtmlBuffer) Code(attrs Attrs, innerHtml func()) {
	b.WriteElement("code", attrs, innerHtml)
}

func (b *HtmlBuffer) Var(attrs Attrs, innerHtml func()) {
	b.WriteElement("var", attrs, innerHtml)
}

func (b *HtmlBuffer) Samp(attrs Attrs, innerHtml func()) {
	b.WriteElement("samp", attrs, innerHtml)
}

func (b *HtmlBuffer) Kbd(attrs Attrs, innerHtml func()) {
	b.WriteElement("kbd", attrs, innerHtml)
}

func (b *HtmlBuffer) Sub(attrs Attrs, innerHtml func()) {
	b.WriteElement("sub", attrs, innerHtml)
}

func (b *HtmlBuffer) Sup(attrs Attrs, innerHtml func()) {
	b.WriteElement("sup", attrs, innerHtml)
}

func (b *HtmlBuffer) I(attrs Attrs, innerHtml func()) {
	b.WriteElement("i", attrs, innerHtml)
}

func (b *HtmlBuffer) B(attrs Attrs, innerHtml func()) {
	b.WriteElement("b", attrs, innerHtml)
}

func (b *HtmlBuffer) U(attrs Attrs, innerHtml func()) {
	b.WriteElement("u", attrs, innerHtml)
}

func (b *HtmlBuffer) Mark(attrs Attrs, innerHtml func()) {
	b.WriteElement("mark", attrs, innerHtml)
}

func (b *HtmlBuffer) Ruby(attrs Attrs, innerHtml func()) {
	b.WriteElement("ruby", attrs, innerHtml)
}

func (b *HtmlBuffer) Rt(attrs Attrs, innerHtml func()) {
	b.WriteElement("rt", attrs, innerHtml)
}

func (b *HtmlBuffer) Rp(attrs Attrs, innerHtml func()) {
	b.WriteElement("rp", attrs, innerHtml)
}

func (b *HtmlBuffer) Bdi(attrs Attrs, innerHtml func()) {
	b.WriteElement("bdi", attrs, innerHtml)
}

func (b *HtmlBuffer) Bdo(attrs Attrs, innerHtml func()) {
	b.WriteElement("bdo", attrs, innerHtml)
}

func (b *HtmlBuffer) Span(attrs Attrs, innerHtml func()) {
	b.WriteElement("span", attrs, innerHtml)
}

func (b *HtmlBuffer) Ins(attrs Attrs, innerHtml func()) {
	b.WriteElement("ins", attrs, innerHtml)
}

func (b *HtmlBuffer) Del(attrs Attrs, innerHtml func()) {
	b.WriteElement("del", attrs, innerHtml)
}

func (b *HtmlBuffer) Iframe(attrs Attrs, innerHtml func()) {
	b.WriteElement("iframe", attrs, innerHtml)
}

func (b *HtmlBuffer) Object(attrs Attrs, innerHtml func()) {
	b.WriteElement("object", attrs, innerHtml)
}

func (b *HtmlBuffer) Video(attrs Attrs, innerHtml func()) {
	b.WriteElement("video", attrs, innerHtml)
}

func (b *HtmlBuffer) Audio(attrs Attrs, innerHtml func()) {
	b.WriteElement("audio", attrs, innerHtml)
}

func (b *HtmlBuffer) Canvas(attrs Attrs, innerHtml func()) {
	b.WriteElement("canvas", attrs, innerHtml)
}

func (b *HtmlBuffer) Map(attrs Attrs, innerHtml func()) {
	b.WriteElement("map", attrs, innerHtml)
}

func (b *HtmlBuffer) Svg(attrs Attrs, innerHtml func()) {
	b.WriteElement("svg", attrs, innerHtml)
}

func (b *HtmlBuffer) Math(attrs Attrs, innerHtml func()) {
	b.WriteElement("math", attrs, innerHtml)
}

func (b *HtmlBuffer) Table(attrs Attrs, innerHtml func()) {
	b.WriteElement("table", attrs, innerHtml)
}

func (b *HtmlBuffer) Caption(attrs Attrs, innerHtml func()) {
	b.WriteElement("caption", attrs, innerHtml)
}

func (b *HtmlBuffer) Colgroup(attrs Attrs, innerHtml func()) {
	b.WriteElement("colgroup", attrs, innerHtml)
}

func (b *HtmlBuffer) Tbody(attrs Attrs, innerHtml func()) {
	b.WriteElement("tbody", attrs, innerHtml)
}

func (b *HtmlBuffer) Thead(attrs Attrs, innerHtml func()) {
	b.WriteElement("thead", attrs, innerHtml)
}

func (b *HtmlBuffer) Tfoot(attrs Attrs, innerHtml func()) {
	b.WriteElement("tfoot", attrs, innerHtml)
}

func (b *HtmlBuffer) Tr(attrs Attrs, innerHtml func()) {
	b.WriteElement("tr", attrs, innerHtml)
}

func (b *HtmlBuffer) Td(attrs Attrs, innerHtml func()) {
	b.WriteElement("td", attrs, innerHtml)
}

func (b *HtmlBuffer) Th(attrs Attrs, innerHtml func()) {
	b.WriteElement("th", attrs, innerHtml)
}

func (b *HtmlBuffer) Form(attrs Attrs, innerHtml func()) {
	b.WriteElement("form", attrs, innerHtml)
}

func (b *HtmlBuffer) Fieldset(attrs Attrs, innerHtml func()) {
	b.WriteElement("fieldset", attrs, innerHtml)
}

func (b *HtmlBuffer) Legend(attrs Attrs, innerHtml func()) {
	b.WriteElement("legend", attrs, innerHtml)
}

func (b *HtmlBuffer) Label(attrs Attrs, innerHtml func()) {
	b.WriteElement("label", attrs, innerHtml)
}

func (b *HtmlBuffer) Button(attrs Attrs, innerHtml func()) {
	b.WriteElement("button", attrs, innerHtml)
}

func (b *HtmlBuffer) Select(attrs Attrs, innerHtml func()) {
	b.WriteElement("select", attrs, innerHtml)
}

func (b *HtmlBuffer) Datalist(attrs Attrs, innerHtml func()) {
	b.WriteElement("datalist", attrs, innerHtml)
}

func (b *HtmlBuffer) Optgroup(attrs Attrs, innerHtml func()) {
	b.WriteElement("optgroup", attrs, innerHtml)
}

func (b *HtmlBuffer) Option(attrs Attrs, innerHtml func()) {
	b.WriteElement("option", attrs, innerHtml)
}

func (b *HtmlBuffer) Textarea(attrs Attrs, innerHtml func()) {
	b.WriteElement("textarea", attrs, innerHtml)
}

func (b *HtmlBuffer) Output(attrs Attrs, innerHtml func()) {
	b.WriteElement("output", attrs, innerHtml)
}

func (b *HtmlBuffer) Progress(attrs Attrs, innerHtml func()) {
	b.WriteElement("progress", attrs, innerHtml)
}

func (b *HtmlBuffer) Meter(attrs Attrs, innerHtml func()) {
	b.WriteElement("meter", attrs, innerHtml)
}

func (b *HtmlBuffer) Details(attrs Attrs, innerHtml func()) {
	b.WriteElement("details", attrs, innerHtml)
}

func (b *HtmlBuffer) Summary(attrs Attrs, innerHtml func()) {
	b.WriteElement("summary", attrs, innerHtml)
}

func (b *HtmlBuffer) Menu(attrs Attrs, innerHtml func()) {
	b.WriteElement("menu", attrs, innerHtml)
}

