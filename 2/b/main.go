package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type solver struct{
  m map[string]int
}

func main() {
  f := "./2/input"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  res := 0
  for scanner.Scan() {
    str := strings.Split(scanner.Text(), ":") 
    s := new(solver)
    s.m = map[string]int{
      "red": 0,
      "green": 0,
      "blue": 0,
    }
    s.analyzeParts(strings.Split(str[1], ";"))
    curRes := s.m["red"] * s.m["green"] * s.m["blue"]
    res += curRes
  }
  fmt.Println("res ", res)
}

func (s *solver) analyzeParts(parts []string) {
  for _, part := range(parts) {
    nums := strings.Split(part, ",")
    s.updateMin(nums)
  }
}

func (s *solver) updateMin(nums []string) {
  for _, els := range(nums) {
    el := strings.Split(els, " ")
    color := el[2]
    val, err := strconv.Atoi(el[1])
    if err != nil {
      fmt.Println(err)
      return
    }
    if s.m[color] < val {
      s.m[color] = val
    }
  }
}