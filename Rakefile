require "yaml"

def notice
  "// This file was generated by the 'generate' Rake task\n\n"
end

def package
  "package htmlbuffer\n\n"
end

def write_header(file)
  file << notice
  file << package
end

def create_elements_go_file(elements)
  elements_go_file = File.open("gen_elements.go", "w") do |f|
    write_header(f)

    elements.each do |el|
      func = <<FUNC
func (b *HtmlBuffer) #{el.capitalize}(attrs Attrs, innerHtml func()) {
\tb.WriteElement("#{el}", attrs, innerHtml)
}

FUNC

      f << func
    end
  end
end

def create_void_elements_go_file(elements)
  void_elements_go_file = File.open("gen_void_elements.go", "w") do |f|
    write_header(f)

    elements.each do |el|
      func = <<FUNC
func (b *HtmlBuffer) #{el.capitalize}(attrs Attrs) {
\tb.WriteVoidElement("#{el}", attrs)
}

FUNC

    f << func
    end
  end
end

desc "Generate go file with elements"
task :generate do
  elements = YAML.load_file("elements.yml")
  create_elements_go_file(elements["non-void"])
  create_void_elements_go_file(elements["void"])
end
