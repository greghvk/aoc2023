package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lim1 = 2e14
var lim2 = 4e14

func main() {
  f := "./24/a/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  eqs := [][]int{}
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    eqs = append(eqs, parse(scanner.Text()))
  }
  res := 0
  for i := range(eqs) {
    for j := i+1; j<len(eqs); j++ {
      if cross(eqs[i], eqs[j], i, j) {
        res++
      }
    }
  }
  fmt.Println(res)
}

func parse(s string) []int {
  els := strings.Split(s, " @ ")
  
  res := []int{}
  for _, e := range(els) {
    for _, i := range(strings.Split(e, ", ")) {
      for _, j := range strings.Fields(i) {
        num, err := strconv.Atoi(j)
        if err != nil {
          panic(err)
        }
        res = append(res, num)
      }
    }
  }
  return res
}

// []int: 0: px 1: py 2: pz @ 3: vx 4: vy 5: vz

func cross(a, b []int, i, j int) bool {
  slopeA := float64(a[4]) / float64(a[3])
  slopeB := float64(b[4]) / float64(b[3])

  if slopeA == slopeB {
    return false
  }
  // eq a: 
  // ya = (x - a[0])*slopeA + a[1]
  // yb = (x - b[0])*slopeB + b[1]
  x := (-slopeB * float64(b[0]) + float64(b[1]) + slopeA * float64(a[0]) - float64(a[1])) / (slopeA - slopeB)
  if !(x >= lim1 && x <= lim2) {
    return false
  }
  y := (x - float64(a[0])) * slopeA + float64(a[1])

  if wrongDir(x, y, a) || wrongDir(x, y, b) {
    return false
  }

  return y >= lim1 && y <= lim2 
}

func wrongDir(x, y float64, v []int) bool {
  if v[3] < 0 && x > float64(v[0]) {
    return true
  }

  if v[3] > 0 && x < float64(v[0]) {
    return true
  }

  if v[4] < 0 && y > float64(v[1]) {
    return true
  }

  if v[4] > 0 && y < float64(v[1]) {
    return true
  }
  return false

}