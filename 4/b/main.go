package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f := "./4/b/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  res := 0

  var counts []int


  i := 0
  for scanner.Scan() {
    if i == len(counts) {
      counts = append(counts, 1)
    } else {
      counts[i]++
    }
    res += counts[i]
    // fmt.Println("card ", i, " has ", counts[i])
    numMatches := solve(scanner.Text())
    // fmt.Println("num matches is ", numMatches)
    for j := 1; j<=numMatches; j++ {
      if i+j == len(counts) {
        counts = append(counts, 0)
      }
      counts[i+j] += counts[i]
    }
    i++
  }
  fmt.Println(res)
}

func solve(s string) int {
  s = strings.Split(s, ":")[1]
  split := strings.Split(s, "|")
  m := makeMap(parseNums(split[0]))
  return count(parseNums(split[1]), m)
}

func parseNums(s string) []int {
  nums := strings.Split(s, " ")
  var res []int
  for _, el := range(nums) {
    i, err := strconv.Atoi(el)
    if err != nil {
      continue
    }
    res = append(res, i)
  }
  return res
}

func makeMap(nums []int) map[int]bool {
  m := make(map[int]bool)
  for _, el := range(nums) {
    m[el] = true
  }
  return m
}

func count(nums []int, m map[int]bool) int {
  res := 0

  for _, el := range(nums) {
    if _, ok := m[el]; ok {
      res++
    }
  }

  return res
}