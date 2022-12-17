package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

type CraneInstruction struct {
  number int
  source int
  destination int
}

func parseInstructions(lines []string) ([]CraneInstruction) {
  instructions := []CraneInstruction{}

  for _, line := range lines {
    var inst CraneInstruction
    _, err := fmt.Sscanf(line, "move %d from %d to %d", &inst.number, &inst.source, &inst.destination)
    if err == nil {
      inst.source--
      inst.destination--
      instructions = append(instructions, inst)
    }
  }

  return instructions
}

func parseLines() []string {
  scanner := bufio.NewScanner(os.Stdin)
  var lines []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  err := scanner.Err()
  if err != nil {
    log.Fatal(err)
  }

  return lines
}

func parseStacks(lines []string) ([][]string) {
  var stacks [][]string

  for i := 0; strings.Contains(lines[i], "["); i++ {
    line := lines[i]
    for j := 1; j < len(line); j += 4 {
      k := (j - 1) / 4
      if len(stacks) == k {
        stacks = append(stacks, []string{})
      }
      if line[j] != ' ' {
        stacks[k] = append([]string{string(line[j])}, stacks[k]...)
      }
    }
  }

  return stacks
}

func main() {
  lines := parseLines()
  stacks := parseStacks(lines)
  instructions := parseInstructions(lines)

  for _, inst := range instructions {
    var x string
    x, stacks[inst.source] = stacks[inst.source][len(stacks[inst.source]) - 1], stacks[inst.source][:len(stacks[inst.source]) - 1]
    stacks[inst.destination] = append(stacks[inst.destination], x)
  }

  fmt.Println(stacks)
}
