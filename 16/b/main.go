package main

import (
	"bufio"
	"fmt"
	"os"
)


var dirs = [][]int{
  {0, 1}, {-1, 0}, {0,-1}, {1, 0},
}

func main() {
  f := "./16/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  

  g := [][]rune{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    g = append(g, []rune(scanner.Text()))
  }

  lastRow := len(g)-1
  lastCol := len(g[0])-1

  res := 0
  var startPoses [][]int
  // startPoses = append(startPoses, []int{0, 0, 3})
  for i := range(g) {
    startPoses = append(startPoses, []int{i, 0, 0}, []int{i, lastCol, 2})
  }
  for j := range(g[0]) {
    startPoses = append(startPoses, []int{0, j, 3}, []int{lastRow, j, 1})
  }
  for i := range(startPoses) {
    cand := solve(g, startPoses[i])
    if cand > res {
      res = cand
    }
  }
  fmt.Println(res)
}

func contains(v []int, i int) bool {
  for _, el := range v {
    if el == i {
      return true
    }
  }
  return false
}

func solve(g [][]rune, startDir []int) int {
  dp := make([][][]int, len(g))
  for i := range(dp) {
    dp[i] = make([][]int, len(g[0]))
  }
  q := [][]int{startDir}

  for len(q) > 0 {
    sz := len(q)

    for sz > 0 {
      sz--
      el := q[0]
      q = q[1:]
      i := el[0]
      j := el[1]
      if i < 0 || i == len(g) || j < 0 || j == len(g[0]) {
        continue
      }
      dir := el[2]
      if contains(dp[i][j], dir) {
        continue
      }
      dp[i][j] = append(dp[i][j], dir)
      q = append(q, getNewDirs(g[i][j], i, j, dir)...)
    }
  }
  
  res := 0
  for i := range(dp) {
    for j := range(dp[i]) {
      if len(dp[i][j]) > 0 {
        res++
      }
    }
  }
  return res

}

func getNewDirs(char rune, i, j, dir int) [][]int {
  switch char {
  case '.':
    return [][]int{moveNormally(i, j, dir)}
  case '|':
    if dir == 1 || dir == 3 {
      return [][]int{moveNormally(i, j, dir)}
    }
    return [][]int{{i+1, j, 3}, {i-1, j, 1}}
  case '-':
    if dir == 0 || dir == 2 {
      return [][]int{moveNormally(i, j, dir)}
    }
    return [][]int{{i, j+1, 0}, {i, j-1, 2}}
  case '/':
    return [][]int{moveNormally(i, j, slashMap[dir])}
  case '\\':
    return [][]int{moveNormally(i, j, backslashMap[dir])}
  default:
    panic("unrecognized direction")
  }
}

func moveNormally(i, j, dir int) []int {
  v := dirs[dir]
  i += v[0]
  j += v[1]
  return []int{i, j, dir}
}

// /
var slashMap = map[int]int{
  0: 1,
  1: 0,
  3: 2,
  2: 3,
}
// \
var backslashMap = map[int]int{
  0: 3,
  3: 0,
  1: 2,
  2: 1,
}
