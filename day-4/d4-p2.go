package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func rangeToBytes(start, end int) ([]byte) {
  var bytes []byte
  for start <= end {
    bytes = append(bytes, byte(start))
    start++
  }
  return bytes
}

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
    var first_start, first_end, second_start, second_end int

    _, err := fmt.Sscanf(line, "%d-%d,%d-%d", &first_start, &first_end, &second_start, &second_end)
    if err != nil {
      panic(err)
    }

    first := rangeToBytes(first_start, first_end)
    second := rangeToBytes(second_start, second_end)

    if strings.ContainsAny(string(first), string(second)) {
      sum++
    }

  }

  fmt.Println(sum)
}
