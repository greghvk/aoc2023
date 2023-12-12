package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var distMultiplier = 1000000

func main() {
  f := "./11/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  row := 0
  var els [][]int
  for scanner.Scan() {
    candidates := getPositions(scanner.Text(), row)
    if len(candidates) == 0 {
      row += distMultiplier
    } else {
      row++
    }
    els = append(els, candidates...)
  }
  sort.Slice(els, func(i, j int) bool {
    return els[i][1] < els[j][1]
  })
  prevCol := 0
  missingCols := 0
  for i:=0; i<len(els); i++ {
    if els[i][1] > prevCol+1 {
      missingCols += (els[i][1] - prevCol - 1) * distMultiplier - 1
    }
    prevCol = els[i][1]
    els[i][1] += missingCols
  }
  var res int64
  for i := range(els) {
    for j := i+1; j<len(els); j++ {
      res += abs(els[i][0] - els[j][0]) + abs(els[i][1] - els[j][1])
      // fmt.Println("dist between ", i , " and ", j, " is ", abs(els[i][0] - els[j][0]) + abs(els[i][1] - els[j][1]))
    }
  }
  fmt.Println(res)
}

func getPositions(s string, row int) [][]int {
  var res [][]int
  for i, c:=range(s) {
    if c == '#' {
      res = append(res, []int{row, i})
    }
  }
  return res
}

func abs(i int) int64 {
  if i < 0 {
    return int64(-i)
  }
  return int64(i)
}