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
  g := [][]rune{}
  for scanner.Scan() {
    g = append(g, []rune(scanner.Text()))
  }
  st := getSt(g)
  q := [][]int{st}
  steps := 0

  res := 0
  for steps <= 64 {
    sz := len(q)
    for sz > 0 {
      sz--
      el := q[0]
      q = q[1:]
      i := el[0]
      j := el[1]
      if steps % 2 == 0 {
        g[i][j] = 'X'
        res++
      }
      for _, dir := range(dirs) {
        newi := i+dir[0]
        newj := j+dir[1]
        if newi<0 || newi==len(g) || newj<0 || newj==len(g[0]) || g[newi][newj] != '.' {
          continue
        }
        // fmt.Println("can travel to ", newi, newj, "with n steps", steps)
        g[newi][newj] = 'O'
        q = append(q, []int{newi, newj})
      }
    }
    steps++
    // fmt.Println("AFTER STEPS", steps)
    // for i := range(g) {
    //   fmt.Println(string(g[i]))
    // }
  }
  fmt.Println(len(g), len(g[0]))
  fmt.Println(res)

}

func getSt(g [][]rune) []int {
  for i := range(g) {
    for j:=range(g[i]) {
      if g[i][j] == 'S' {
        return []int{i, j}
      }
    }
  }
  return nil
}