package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  f := "./14/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  var r [][]rune
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    r = append(r, []rune(scanner.Text()))
  }
  fmt.Println(solve(r))
}

func solve(r [][]rune) int {
  moveNorth(r)

  res := 0
  for i := range(r) {
    for j := range(r[i]) {
      if r[i][j] != 'O' {
        continue
      }
      res += len(r) - i
    }
  }
  return res
}

func moveNorth(r [][]rune) {
  for j := range(r[0]) {
    pos := len(r)-1
    for i := len(r)-1; i>=0; i-- {
      if r[i][j] == '#' {
        pos = i-1
        continue
      }
      if r[i][j] == '.' {
        r[i][j] = r[pos][j]
        r[pos][j] = '.'
        pos--
      }
    }
  }

}