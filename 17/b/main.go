package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var dirs = [][]int{
  {0, 1}, {-1, 0}, {0,-1}, {1, 0},
}


type Heap [][]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
  f := "./17/b/input"
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
  dp := make([][][][]int,len(r))
  for i := range(r) {
    dp[i] = make([][][]int, len(r[i]))
  }
  h := &Heap{{0, 0, 0, 0, 0}, {0, 0, 0, 3, 0}}
  for len(*h) > 0 {
    el := heap.Pop(h).([]int)
    heatLoss := el[0]
    i := el[1]
    j := el[2]
    dir := el[3]
    length := el[4]
    if i == len(dp)-1 && j == len(dp[0])-1 && length >= 4 {
      fmt.Println(heatLoss)
      return
    }
    if contains(dp[i][j], []int{dir, length}) {
      continue
    }
    dp[i][j] = append(dp[i][j], []int{dir, length})
    candidates := getCandidates(i, j, dir, length, len(dp), len(dp[0]))
    for _, cand := range(candidates) {
      newi := cand[0]
      newj := cand[1]
      newDir := cand[2]
      newLen := cand[3]
      newLoss := heatLoss + int(r[newi][newj]) - int('0')
      heap.Push(h, []int{newLoss, newi, newj, newDir, newLen})
    }
  }
  fmt.Println("donew without res")
}

func getCandidates(i, j, dir, length, m, n int) [][]int {
  res := [][]int{}
  for k := -1; k<2; k++ {
    if (k == -1 || k==1) && length < 4 {
      continue
    }
    newDir := dir + k
    if newDir < 0 {
      newDir = 3
    }
    newDir %= 4
    newLength := 1
    if k==0 {
      if length==10 {
        continue
      }
      newLength = length+1
    }
    newPos := moveDir(i, j, newDir, m, n)
    if len(newPos) == 0 {
      continue
    }
    res = append(res, []int{newPos[0], newPos[1], newDir, newLength})


  }
  return res
}

func moveDir(i, j, dir, m, n int) []int {
  v := dirs[dir]
  newi := i + v[0]
  newj := j + v[1]

  if newi<0 || newi==m || newj<0 || newj==n {
    return nil
  }
  return []int{newi, newj}
}

func contains(v [][]int, el []int) bool {
  for i := range(v) {
    ok := true
    for j := range(v[i]) {
      if v[i][j] != el[j] {
        ok = false
        break
      }
    }
    if ok {
      return true
    }
  }
  return false
}