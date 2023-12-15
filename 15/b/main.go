package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  f := "./15/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()
  txt := strings.Split(scanner.Text(), ",")
  res := 0
  for i := range(txt) {
    res += solve([]rune(txt[i]))
  }
  fmt.Println(res)
}

func solve(r []rune) int {
  res := 0
  for _, c := range r {
    res += int(c)
    res *= 17
    res %= 256
  }
  return res
}