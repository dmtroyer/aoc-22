#!/usr/bin/env ruby

alphabet = [*'a'..'z', *'A'..'Z']
sum = 0

File.foreach('input.txt', chomp: true) do |line|
  first, second = line.chars.each_slice(line.length / 2).to_a
  item = first.intersection(second).first
  sum += alphabet.index(item) + 1
end

puts sum
