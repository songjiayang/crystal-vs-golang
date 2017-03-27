require "benchmark"

f = (File.expand_path("~") + "/crystal/src/class.cr")

Benchmark.bm do |x|
  x.report("") { 500000.times do
    File.read(f).lines.size
  end }
end
