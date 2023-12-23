package main

import (
	"bufio"
	"fmt"
	"os"
)

var dirs = [][]int{
  {-1, 0, int('^')},
  {1, 0, int('v')},
  {0, 1, int('>')},
  {0, -1, int('<')},
}

func main() {
  f := "./23/a/input"
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

  dp := make([][]int, len(r))
  for i := range(dp) {
    dp[i] = make([]int, len(r[0]))
    for j := range(dp[i]) {
      dp[i][j] = -1
    }
  }

  res := solve(cur, r, mem, dp) - 1 // not counting the startig tile
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

func solve(cur []int, r [][]rune, mem [][]bool, dp [][]int) int {
  if dp[cur[0]][cur[1]] != -1 {
    return dp[cur[0]][cur[1]]
  }
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
    var cand int

    if r[newi][newj] == '.' {
      cand = solve([]int{newi, newj}, r, mem, dp)
    } else if int(r[newi][newj]) == dir[2] {
      newi2 := newi+dir[0]
      newj2 := newj+dir[1]
      if !mem[newi2][newj2] {
        cand = 1 + solve([]int{newi2, newj2}, r, mem, dp)
      }
    }
    if cand > maxRes {
      maxRes = cand
    }
  }
  mem[cur[0]][cur[1]] = false
  if maxRes == 0 {
    dp[cur[0]][cur[1]] = 0
    return 0
  }

  dp[cur[0]][cur[1]] = maxRes+1
  return maxRes+1

}
