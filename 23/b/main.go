package main

import (
	"bufio"
	"fmt"
	"os"
)


var dirs = [][]int{
  {-1, 0},
  {1, 0},
  {0, 1},
  {0, -1},
}

func main() {
  f := "./23/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  r := [][]rune{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    r = append(r, []rune(scanner.Text()))
  }

  cur := getStart(r[0])
  mem := make([][]bool, len(r))
  for i := range(mem) {
    mem[i] = make([]bool, len(r[0]))
  }

  res := solve(cur, r, mem) - 1 // not counting the startig tile
  fmt.Println(res)
}

func getStart(r []rune) []int {
  for i := range(r) {
    if r[i] == '.' {
      return []int{0, i}
    }
  }
  return nil
}

func solve(cur []int, r [][]rune, mem [][]bool) int {
  if cur[0] == len(r)-1 {
    return 1
  }
  mem[cur[0]][cur[1]] = true
  maxRes := 0
  for _, dir := range(dirs) {
    newi := cur[0] + dir[0]
    newj := cur[1] + dir[1]
    if newi<0 || newi == len(r) || newj<0 || newj==len(r[0]) || r[newi][newj] == '#' || mem[newi][newj] {
      continue
    }
    cand := solve([]int{newi, newj}, r, mem)
    if cand > maxRes {
      maxRes = cand
    }
  }
  mem[cur[0]][cur[1]] = false
  if maxRes == 0 {
    return 0
  }

  return maxRes+1

}
