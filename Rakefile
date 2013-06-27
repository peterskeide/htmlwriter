require "yaml"
require_relative "generators"

desc "Generate go file with elements"
task :generate do
  elements = YAML.load_file("elements.yml")
  NormalElementGenerator.new.generate(elements["non-void"])
  VoidElementGenerator.new.generate(elements["void"])
  InputElementGenerator.new.generate(elements["inputs"])
  TextOnlyElementGenerator.new.generate(elements["text"])
end
