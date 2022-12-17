#!/usr/bin/env ruby

# (1 for A Rock, 2 for B Paper, and 3 for C Scissors)
# X = lose (0), Y = draw (3), Z = win (6)

perms = {
  "A X" => 3, # lose with scissors
  "A Y" => 4, # draw with rock
  "A Z" => 8, # win with paper
  "B X" => 1, # lose with rock
  "B Y" => 5, # draw with paper
  "B Z" => 9, # win with scissors
  "C X" => 2, # lose with paper
  "C Y" => 6, # draw with scissors
  "C Z" => 7  # win with rock
}

score = 0

File.foreach('input.txt', chomp: true) do |line|
  score += perms[line]
end

puts score
