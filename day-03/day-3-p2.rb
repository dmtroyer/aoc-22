#!/usr/bin/env ruby

GROUP_SIZE = 3

alphabet = [*'a'..'z', *'A'..'Z']
elf = sum = 0
groups = Array.new

File.foreach('input.txt', chomp: true) do |line|
  groups[elf / GROUP_SIZE] ||= Array.new
  groups[elf / GROUP_SIZE][elf % GROUP_SIZE] = line.chars

  group = groups[elf / GROUP_SIZE]
  if group.length == GROUP_SIZE
    intersection = group[0] & group[1] & group[2]
    throw "NOPE." if intersection.length != 1
    sum += alphabet.index(intersection.first) + 1
  end

  elf += 1
end

puts sum
