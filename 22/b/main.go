package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type brick struct {
  dims [][]int
  id, h int
}

func (b brick) height() int {
  return b.dims[2][1] - b.dims[2][0] + 1
}

func (b brick) fields() [][]int {
  st := []int{b.dims[0][0], b.dims[1][0]}
  end := []int{b.dims[0][1], b.dims[1][1]}
  res := [][]int{}
  for {
    res = append(res, []int{st[0], st[1]})
    if st[0] == end[0] && st[1] == end[1] {
      break
    }
    if st[0] < end[0] {
      st[0]++
    }
    if st[1] < end[1] {
      st[1]++
    }
  }
  return res
}

func main() {
  f := "./22/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  bricks := []brick{}
  i := 0
  for scanner.Scan() {
    bricks = append(bricks, getBrick(scanner.Text()))
    bricks[len(bricks)-1].id = i
    i++
  }
  sort.Slice(bricks, func(i, j int) bool {
    return bricks[i].dims[2][0] < bricks[j].dims[2][0]
  })
  fmt.Println(solve(bricks))
}

func getBrick(s string) brick {
  els := strings.Split(s, "~")

  res := brick{}
  res.dims = make([][]int, 3)
  for i, s := range(strings.Split(els[0], ",")) {
    num, _ := strconv.Atoi(s)
    res.dims[i] = append(res.dims[i], num)
  }
  for i, s := range(strings.Split(els[1], ",")) {
    num, _ := strconv.Atoi(s)
    res.dims[i] = append(res.dims[i], num)
    sort.Ints(res.dims[i])
  }
  return res
}

func solve(bricks []brick) int {
  resGraph := make(map[int][]int)
  numSupporting := make(map[int]int)
  r := 0
  highestBrick := make([][]brick, 10)
  for i := range(highestBrick) {
    highestBrick[i] = make([]brick, 10)
    for j := range(highestBrick[i]) {
      highestBrick[i][j] = brick{
        h: 0,
        id: -1,
      }
    }
  }

  triggers := map[int]bool{}
  for i := range(bricks) {
    h := bricks[i].height()
    maxHeight := 0
    m := make(map[int]bool)
    for _, f := range(bricks[i].fields()) {
      cand := highestBrick[f[0]][f[1]].h
      if cand == maxHeight {
        m[highestBrick[f[0]][f[1]].id] = true
      } else if cand>maxHeight {
        m = map[int]bool{
          highestBrick[f[0]][f[1]].id: true,
        }
        maxHeight = cand
      }
    }
    for _, f := range(bricks[i].fields()) {
      highestBrick[f[0]][f[1]] = brick{
        h: maxHeight + h,
        id: bricks[i].id,
      }
    }
    for k := range(m) {
      resGraph[k] = append(resGraph[k], bricks[i].id)
    }
    numSupporting[bricks[i].id] = len(m)
    if len(m) == 1 {
      for k := range(m) {
        if k == -1 {
          break
        }
        triggers[k] = true
      }
    }
  }
  for t := range(triggers) {
    cpy := make(map[int]int)
    for k, v := range numSupporting {
      cpy[k] = v
    }
    rs := traverse(t, resGraph, cpy)-1
    r += rs
  }

  return r
}

func traverse(cur int, resGraph map[int][]int, numSupporting map[int]int) int {
  res := 1
  for _, nbr := range(resGraph[cur]) {
    numSupporting[nbr]--
    if numSupporting[nbr] == 0 {
      res += traverse(nbr, resGraph, numSupporting)
    }
  }
  return res
}
