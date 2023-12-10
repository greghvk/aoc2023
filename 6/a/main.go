package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  f := "./6/a/input_test"
  file, err := os.Open(f)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Scan()
  times := getFields(scanner.Text())
  scanner.Scan()
  distances := getFields(scanner.Text())

  res := 1
  for i := 0; i < len(times); i++ {
    res *= solve(times[i], distances[i])
    if res == 0 {
      break
    }
  }
  fmt.Println("total res is ", res)
}

func getFields(s string) []int {
  var res []int
  for _, el := range(strings.Fields(s)[1:]) {
    i, _ := strconv.Atoi(el)
    res = append(res, i)
  }
  return res
}

func solve(time, dist int) int {
  curSpeed := 0
  res := 0
  for i:=0; i<time; i++ {
    curSpeed++
    timeLeft := time-i-1
    if curSpeed * timeLeft > dist {
      res++
    } else if res > 0 {
      return res
    }
  }
  return res
}