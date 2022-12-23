package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  var lines []string
  for {
    scanner.Scan()
    line := scanner.Text()

    // break the loop if line is empty
    if len(line) == 0 {
      break
    }
    lines = append(lines, line)
  }

  err := scanner.Err()
  if err != nil {
    log.Fatal(err)
  }

  sum := 0

  for _, line := range lines {
    first := make([]int, 2)
    second := make([]int, 2)

    _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &first[0], &first[1], &second[0], &second[1])
    if err != nil {
      panic(err)
    }

    if (first[0] >= second[0] && first[1] <= second[1]) || (second[0] >= first[0] && second[1] <= first[1]) {
      sum++
    }
  }

  fmt.Println(sum)
}
