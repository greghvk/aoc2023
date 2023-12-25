package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [][]int{
  {-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
  f := "./21/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  stg := [][]rune{}
  for scanner.Scan() {
    stg = append(stg, []rune(scanner.Text()))
  }
  st := getSt(stg)
  g := make([][]rune, len(stg))
  stg[st[0]][st[1]] = '.'
  st[0] += 131*4
  st[1] += 131*4
  
  for i := 0; i<12; i++ {
    for j := range(stg) {
      g[j] = append(g[j], stg[j]...)
    }
  }

  for i := 0; i<10; i++ {
    for j := range(stg) {
      g = append(g, []rune{})
      g[len(g)-1] = make([]rune, len(g[0]))
      copy(g[len(g)-1], g[j])
    }
  }
  q := [][]int{{st[0], st[1]}}
  steps := 0

  res := 0
  for steps <= 65 + 5*131 {
    sz := len(q)
    for sz > 0 {
      sz--
      el := q[0]
      q = q[1:]
      i := el[0]
      j := el[1]
      if steps % 2 == 1 {
        g[i][j] = 'X'
        res++
      }
      for _, dir := range(dirs) {
        newi := i+dir[0]
        newj := j+dir[1]
        if newi<0 || newi==len(g) || newj<0 || newj==len(g[0]) || g[newi][newj] != '.' {
          continue
        }
        g[newi][newj] = 'O'
        q = append(q, []int{newi, newj})
      }
    }
    if steps == 65 || (steps-65) % 262 == 0 {
      fmt.Println("res at", steps, res)
    }
    // this gives us 3 points, we can fit a quadratic curve:
    // https://www.wolframalpha.com/input?i=quadratic+fit+calculator&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3x%22%7D+-%3E%22%7B0%2C+2%2C+4%7D%22&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3y%22%7D+-%3E%22%7B3648%2C+90972%2C+294528%7D%22
    // if we solve for x = 202300 ((26501365-65)/131)) - 65 is starting point, 131 is grid width.
    // we use modulo 262 not 131 because we want odd steps, not even
    steps++
  }

}

func getSt(g [][]rune) []int {
  for i := range(g) {
    for j:=range(g[i]) {
      if g[i][j] == 'S' {
        g[i][j] = '.'
        return []int{i, j}
      }
    }
  }
  return nil
}