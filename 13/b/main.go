package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  f := "./13/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()
  
  var cur [][]rune

  scanner := bufio.NewScanner(file)
  res := 0
  for scanner.Scan() {
    txt := scanner.Text()
    if len(txt) == 0 {
      res += solve(cur)
      cur = [][]rune{}
      continue
    }
    cur = append(cur, []rune(txt))
  }
  res += solve(cur)
  fmt.Println(res)


}

func solve(r [][]rune) int {
  for i := 0; i<len(r)-1; i++ {
    if numDiffsRow(r, i) == 1 {
      return (i+1) * 100
    }
  }
  for i := 0; i<len(r[0])-1; i++ {
    if numDiffsColumn(r, i) == 1 {
      return i+1
    }
    
  }
  return -1
}

func numDiffsRow(r [][]rune, row int) int {
  r2 := row+1
  res := 0
  for {
    if r2 == len(r) || row < 0 {
      return res
    }
    for i:= range r[0] {
      if r[row][i] != r[r2][i] {
        res++
      }
    }
    row--
    r2++
  }
  return -1
}

func numDiffsColumn(r [][]rune, col int) int {
  c2 := col+1
  res := 0
  for {
    if col < 0 || c2 == len(r[0]) {
      return res
    }
    for i := range r {
      if r[i][col] != r[i][c2] {
        res++
      }
    }
    col--
    c2++
  }
  return -1
  
}