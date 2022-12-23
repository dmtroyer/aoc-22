#!/usr/bin/env ruby

elves = Array.new
current_elf = 0

File.foreach('elves.txt', chomp: true) do |line|
  if line.empty?
    current_elf += 1
  else
    elves[current_elf] = 0 if elves[current_elf].nil?
    elves[current_elf] += line.to_i
  end
end

elves.sort!.reverse!
puts elves[0, 3].sum
