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
  m := make(map[string]int)
  cyclesLeft := -1

  for i:=0;i<1e9;i++ {
    moveNorth(r)
    moveWest(r)
    moveSouth(r)
    moveEast(r)

    var s string
    for _, row := range(r) {
      s += string(row)
    }
    if v, ok := m[s]; ok {
      cyclesLeft = (1e9-v-1) % (i-v)
      break
    } else {
      m[s] = i
    }
  }
  for i:=0;i<cyclesLeft;i++ {
    moveNorth(r)
    moveWest(r)
    moveSouth(r)
    moveEast(r)

  }

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

func moveWest(r [][]rune) {
  for i := range(r) {
    pos := len(r[0])-1
    for j := len(r[0])-1; j>=0; j-- {
      if r[i][j] == '#' {
        pos = j-1
        continue
      }

      if r[i][j] == '.' {
        r[i][j] = r[i][pos]
        r[i][pos] = '.'
        pos--
      }
    }
  }
}

func moveSouth(r [][]rune) {
  for j := range(r[0]) {
    pos := 0
    for i := range(r) {
      if r[i][j] == '#' {
        pos = i+1
        continue
      }
      if r[i][j] == '.' {
        r[i][j] = r[pos][j]
        r[pos][j] = '.'
        pos++
      }
    }
  }
}


func moveEast(r [][]rune) {
  for i := range(r) {
    pos := 0
    for j := range(r[0]) {
      if r[i][j] == '#' {
        pos = j+1
        continue
      }

      if r[i][j] == '.' {
        r[i][j] = r[i][pos]
        r[i][pos] = '.'
        pos++
      }
    }
  }
}