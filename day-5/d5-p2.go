package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

type CraneInstruction struct {
  num_crates int
  src_stack int
  dest_stack int
}

func parseInstructions(lines []string) ([]CraneInstruction) {
  instructions := []CraneInstruction{}

  for _, line := range lines {
    var inst CraneInstruction
    _, err := fmt.Sscanf(line, "move %d from %d to %d", &inst.num_crates, &inst.src_stack, &inst.dest_stack)
    if err == nil {
      inst.src_stack--
      inst.dest_stack--
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

func parseStacks(lines []string) ([][]byte) {
  var stacks [][]byte

  for i := 0; strings.Contains(lines[i], "["); i++ {
    line := lines[i]
    for j := 1; j < len(line); j += 4 {
      k := (j - 1) / 4
      if len(stacks) == k {
        stacks = append(stacks, []byte{})
      }
      if line[j] != ' ' {
        line_as_slice := []byte{line[j]}
        stacks[k] = append(line_as_slice, stacks[k]...)
      }
    }
  }

  return stacks
}

func main() {
  lines := parseLines()
  stacks := parseStacks(lines)
  instructions := parseInstructions(lines)

  for _, i := range instructions {
    src := &stacks[i.src_stack]
    dest := &stacks[i.dest_stack]

    bottom_crate := len(*src) - i.num_crates
    moved := (*src)[bottom_crate:]
    *src = (*src)[:bottom_crate]
    *dest = append(*dest, moved...)
  }

  var answer string
  for _, stack := range stacks {
    answer += string(stack[len(stack)-1])
  }
  fmt.Println(answer)
}
