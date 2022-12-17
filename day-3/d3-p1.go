package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

func commonChar(a, b string) byte {
  for _, c := range a {
    if strings.Contains(b, string(c)) {
      return byte(c)
    }
  }
  return 0
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

  alphabet := []byte{}
  var i byte
  for i = 'a'; i <= 'z'; i++ {
    alphabet = append(alphabet, i)
  }
  for i = 'A'; i <= 'Z'; i++ {
    alphabet = append(alphabet, i)
  }

  sum := 0
  for _, line := range lines {
    first := line[:len(line)/2]
    second := line[len(line)/2:]
    common := commonChar(first, second)
    sum += strings.IndexByte(string(alphabet), common) + 1
  }

  fmt.Println(sum)

}
