channel = Channel(Int32).new
n = 100_000

n.times do |_|
  spawn do
  	sleep 0.001 # io block 1ms
    channel.send 1
  end
end

total = 0
n.times do
  total += channel.receive
end
